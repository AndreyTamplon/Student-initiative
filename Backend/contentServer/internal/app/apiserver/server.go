package apiserver

import (
	"APIServerSI/internal/app/apiserver/authentification"
	"APIServerSI/model/petition"
	"APIServerSI/store"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type server struct {
	router    *mux.Router
	logger    *logrus.Logger
	store     store.Store
	jwtConfig *authentification.JWTConfig
}

const (
	ctxUserId       ctxKey = iota
	ctxRequestIDKey ctxKey = iota
)

type ctxKey int16

func NewServer(store store.Store, jwtConfig *authentification.JWTConfig) *server {
	server := &server{
		router:    mux.NewRouter(),
		logger:    logrus.New(),
		store:     store,
		jwtConfig: jwtConfig,
	}

	server.configureRouter()
	return server
}

func (server *server) configureRouter() {
	server.router.Use(server.SetRequestID)
	server.router.Use(server.LogRequest)
	server.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	server.router.HandleFunc("/petitions", server.HandlePetitionsList()).Methods("GET")
	private := server.router.PathPrefix("/api").Subrouter()
	private.Use(server.authenticateUser)
	private.HandleFunc("/create-petition", server.HandlePetitionCreate()).Methods("POST")
	private.HandleFunc("/sign-petition", server.HandleSignPetition()).Methods("POST")
}

func (server *server) SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxRequestIDKey, id)))

	})
}

func (server *server) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := server.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxRequestIDKey),
		})
		logger.Infof("started %s %s", r.Method, r.URL.Path)
		start := time.Now()
		rw := &ResponseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)

	})
}

func (server *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.logger.Infof("authenticating user")
		token := r.Header.Get("token")
		if token == "" {
			server.error(w, r, http.StatusUnauthorized, ErrNoToken)
			return
		}
		id, err := server.jwtConfig.ParseToken(token)
		if err != nil {
			server.error(w, r, http.StatusUnauthorized, ErrInvalidToken)
			server.logger.Errorf("failed to parse token: %v", err)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserId, id)))
	})
}

func (server *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func (server *server) HandlePetitionCreate() http.HandlerFunc {
	type request struct {
		Title            string   `json:"title"`
		AuthorName       string   `json:"authorName"`
		AuthorEmail      string   `json:"authorEmail"`
		DateOfCreation   string   `json:"dateOfCreation"`
		DateOfExpiration string   `json:"dateOfExpiration"`
		Tags             []string `json:"tags"`
		PetitionContent  string   `json:"petitionContent"`
		SignaturesTarget int      `json:"signaturesTarget"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Info("HandlePetitionCreate")
		request := request{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			server.logger.Errorf("failed to decode request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		petition := petition.Petition{
			Title:              request.Title,
			AuthorName:         request.AuthorName,
			AuthorEmail:        request.AuthorEmail,
			DateOfCreation:     request.DateOfCreation,
			DateOfExpiration:   request.DateOfExpiration,
			Tags:               request.Tags,
			PetitionContent:    request.PetitionContent,
			NumberOfSignatures: 0,
			SignaturesTarget:   request.SignaturesTarget,
		}
		if _, err := server.store.Petitions().FindByTitle(petition.Title); err == nil {
			server.logger.Infof("petition with title %s already exists", petition.Title)
			server.error(w, r, http.StatusConflict, ErrPetitionAlreadyExists)
			return
		} else if err != store.ErrRecordNotFound {
			server.logger.Errorf("failed to find petition by title: %v", err)
			server.error(w, r, http.StatusInternalServerError, err)
			return
		}
		if err := server.store.Petitions().Create(&petition); err != nil {
			server.logger.Errorf("failed to create petition: %v", err)
			server.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := server.store.Petitions().SignPetition(r.Context().Value(ctxUserId).(int), petition.ID); err != nil {
			server.logger.Errorf("failed to sign petition: %v", err)
			server.error(w, r, http.StatusInternalServerError, err)
			return
		}
		server.logger.Infof("petition %s successfully created", petition.Title)
		server.respond(w, r, http.StatusOK, nil)
	}
}
func (server *server) HandlePetitionsList() http.HandlerFunc {
	type request struct {
		IdFrom int `json:"idFrom"`
		IdTo   int `json:"idTo"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Info("HandlePetitionList")
		request := request{}
		var err error
		request.IdFrom, err = strconv.Atoi(r.Header.Get("idFrom"))
		if err != nil {
			server.logger.Errorf("failed to parse idFrom or idTo: %v", err)
			server.error(w, r, http.StatusBadRequest, err)
			return
		}
		request.IdTo, err = strconv.Atoi(r.Header.Get("idTo"))
		if err != nil {
			server.logger.Errorf("failed to parse idFrom or idTo: %v", err)
			server.error(w, r, http.StatusBadRequest, err)
			return
		}
		petitions, err := server.store.Petitions().GetPetitionsByIdRange(request.IdFrom, request.IdTo)
		if err != nil {
			server.logger.Infof("failed to get petitions by id range: %v", err)
			server.error(w, r, http.StatusInternalServerError, err)
			return
		}
		server.logger.Infof("petitions list successfully sent")
		server.respond(w, r, http.StatusOK, petitions)
	}
}

func (server *server) HandleSignPetition() http.HandlerFunc {
	type request struct {
		PetitionId int    `json:"petitionId"`
		UserEmail  string `json:"userEmail"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Info("HandleSignPetition")
		request := request{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			server.logger.Errorf("failed to decode request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := server.store.Petitions().SignPetition(r.Context().Value(ctxUserId).(int), request.PetitionId); err != nil {
			if err == store.ErrAlreadySigned {
				server.logger.Infof("user %s already signed petition %d", request.UserEmail, request.PetitionId)
				server.error(w, r, http.StatusConflict, ErrAlreadySigned)
				return
			} else {
				server.logger.Errorf("failed to sign petition: %v", err)
				server.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}
		server.logger.Infof("petition %s successfully signed", request.PetitionId)
		server.respond(w, r, http.StatusOK, nil)
	}
}
func (server *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	server.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (server *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
