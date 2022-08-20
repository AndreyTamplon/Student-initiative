package main

import (
	"APIServerSI/internal/app/apiserver"
	"fmt"
	"log"
	"os"
)

func main() {
	dataBaseURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	serverConfig := &apiserver.Config{
		BindAddr:    os.Getenv("BIND_ADDR"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
		DatabaseURL: dataBaseURL,
		JwtSecret:   os.Getenv("JWT_SECRET"),
		JwtTTL:      os.Getenv("JWT_TTL"),
	}
	if err := apiserver.Start(serverConfig); err != nil {
		log.Fatal("failed to start server", err)
	}
}
