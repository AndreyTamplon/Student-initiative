package apiserver

import (
	"authorizationServer/internal/app/apiserver/authentification"
	"authorizationServer/model/user"
	"authorizationServer/store"
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
	router      *mux.Router
	logger      *logrus.Logger
	store       store.Store
	EmailSender *authentification.EmailSender
	jwtConfig   *authentification.JWTConfig
}

const (
	ctxRequestIDKey ctxKey = iota
)

type ctxKey int16

func NewServer(store store.Store, sender *authentification.EmailSender, jwtConfig *authentification.JWTConfig) *server {
	server := &server{
		router:      mux.NewRouter(),
		logger:      logrus.New(),
		store:       store,
		EmailSender: sender,
		jwtConfig:   jwtConfig,
	}

	server.configureRouter()
	return server
}

func (server *server) configureRouter() {
	server.router.Use(server.SetRequestID)
	server.router.Use(server.LogRequest)
	server.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	server.router.HandleFunc("/sign-up", server.HandleUsersCreate()).Methods("POST")
	server.router.HandleFunc("/sign-up-confirmation", server.HandleRegistrationConfirmation()).Methods("POST")
	server.router.HandleFunc("/sign-in", server.HandleLogin()).Methods("POST")
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

func (server *server) HandleLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Info("HandleLogin")
		request := request{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			server.logger.Errorf("failed to decode request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := server.store.Users().FindByEmail(request.Email)
		if err != nil || !user.ComparePassword(request.Password) {
			server.logger.Infof("failed to find user by email: %v", err)
			server.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailOrPassword)
			return
		}
		token, err := server.jwtConfig.GenerateJWT(user.ID)
		if err != nil {
			server.logger.Errorf("failed to generate token: %v", err)
			server.error(w, r, http.StatusInternalServerError, err)
			return
		}
		server.logger.Infof("user %s successfully logged in", user.Name)
		server.respond(w, r, http.StatusOK, map[string]string{"token": token, "name": user.Name})
	}
}

func (server *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func (server *server) HandleUsersCreate() http.HandlerFunc {
	type request struct {
		Username string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Info("HandleUsersCreate")
		request := request{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			server.logger.Errorf("failed to decode request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		code := authentification.GenerateRandomNumber(100000, 999999)
		encryptedConfirmationCode := authentification.EncryptCode(code)
		user := &user.User{
			Name:                      request.Username,
			Email:                     request.Email,
			Confirmed:                 false,
			EncryptedConfirmationCode: encryptedConfirmationCode,
			Password:                  request.Password,
		}
		userInDB, err := server.store.Users().FindByEmail(user.Email)
		if err == nil && userInDB != nil {
			server.logger.Infof("user with email %s already exists", user.Email)
			server.error(w, r, http.StatusConflict, ErrIncorrectEmailOrPassword)
			return
		}

		if err := server.store.Users().Create(user); err != nil {
			server.logger.Errorf("failed to create user: %v", err)
			server.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		recipientAddress := []*authentification.Address{
			{
				Name:  user.Name,
				Email: user.Email,
			},
		}
		message := "Код подтверждения: " + strconv.Itoa(code)
		if err := server.EmailSender.DefaultSend(message, recipientAddress); err != nil {
			server.logger.Errorf("failed to send email: %v", err)
			server.error(w, r, http.StatusInternalServerError, err)
			return
		}
		user.Sanitize()
		server.logger.Infof("Confirmation code has been send to %s", user.Email)
		server.respond(w, r, http.StatusOK, user)
	}
}

func (server *server) HandleRegistrationConfirmation() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.logger.Info("HandleRegistrationConfirmation")
		request := request{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			server.logger.Errorf("Error decoding request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := server.store.Users().FindByEmail(request.Email)
		if err != nil {
			server.logger.Infof("User not found: %v", err)
			server.error(w, r, http.StatusNotFound, err)
			return
		}
		if authentification.CompareCode(request.Code, user.EncryptedConfirmationCode) {
			user.Confirmed = true
			if err := server.store.Users().UpdateUser(user); err != nil {
				server.logger.Errorf("Error updating user: %v", err)
				server.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			server.logger.Infof("User %s confirmed", user.Email)
			server.respond(w, r, http.StatusCreated, user)
		} else if user.Confirmed {
			server.logger.Infof("User %s already confirmed", user.Email)
			server.error(w, r, http.StatusConflict, ErrAlreadyConfirmed)
			return
		} else {
			server.logger.Infof("User %s incorrect confirmation code", user.Email)
			server.error(w, r, http.StatusUnauthorized, ErrInvalidCode)
		}
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
