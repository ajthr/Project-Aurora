package main

import (
	"log"
	"net/http"
	"os"

	"identity-service/internal/config"
	"identity-service/internal/routers"
)

var (
	// db variables
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")

	// jwt variables
	secretKey = os.Getenv("JWT_SECRET")

	// google auth variables
	googleClientId = os.Getenv("GOOGLE_CLIENT_ID")

	// mail sender variables
	mailFrom     = os.Getenv("SMTP_MAIL_ID")
	mailPassword = os.Getenv("SMTP_PASSWORD")
	mailHost     = os.Getenv("SMTP_HOST")
	mailPort     = os.Getenv("SMTP_PORT")
)

func main() {
	// create and get database connection
	dbConfig, err := config.NewDBConfig(host, port, user, password, dbname)

	if err != nil {
		log.Fatal("ERROR cannot connect to database", err)
	}

	// jwt config
	jwtConfig := config.NewJWTConfig(secretKey)

	// google client config
	googleClient := config.NewGoogleAuthClient(googleClientId)

	// create mail sender
	mailClient := config.NewMailConfig(mailFrom, mailPassword, mailHost, mailPort)

	// create router with given configurations
	router := routers.NewRouter(dbConfig, jwtConfig, googleClient, mailClient)

	log.Println("Ready to receive requests.")
	http.ListenAndServe(":7000", router)
}
