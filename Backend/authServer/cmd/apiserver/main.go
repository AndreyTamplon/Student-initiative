package main

import (
	"authorizationServer/internal/app/apiserver"
	"authorizationServer/internal/app/apiserver/authentification"
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
	emailSenderConfig := &authentification.EmailSenderConfiguration{
		APIKey:       os.Getenv("API_KEY"),
		FromName:     os.Getenv("FROM_NAME"),
		FromEmail:    os.Getenv("FROM_EMAIL"),
		Subject:      os.Getenv("SUBJECT"),
		ReplyToName:  os.Getenv("REPLY_TO_NAME"),
		ReplyToEmail: os.Getenv("REPLY_TO_EMAIL"),
	}
	emailSender := authentification.NewEmailSender(emailSenderConfig)

	if err := apiserver.Start(serverConfig, emailSender); err != nil {
		log.Fatal("failed to start server", err)
	}
}
