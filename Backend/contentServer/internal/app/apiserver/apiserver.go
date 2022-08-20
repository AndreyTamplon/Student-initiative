package apiserver

import (
	"APIServerSI/internal/app/apiserver/authentification"
	"APIServerSI/store/SQLStore"
	"database/sql"
	"github.com/rs/cors"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := SQLStore.NewStore(db)
	jwtConfig, err := authentification.ConfigureJWT(config.JwtSecret, config.JwtTTL)
	if err != nil {
		return err
	}
	server := NewServer(store, jwtConfig)
	c := cors.AllowAll()
	return http.ListenAndServe(config.BindAddr, c.Handler(server))
}

func newDB(DatabaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
