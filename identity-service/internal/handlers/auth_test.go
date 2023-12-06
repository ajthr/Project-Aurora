package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"identity-service/internal/config"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	// jwt variables
	secretKey = os.Getenv("JWT_SECRET")

	// google auth variables
	googleClientId = os.Getenv("GOOGLE_CLIENT_ID")

	// mailtrap smtp test server variables
	mailFrom     = os.Getenv("MAILTRAP_SMTP_USERNAME")
	mailPassword = os.Getenv("MAILTRAP_SMTP_PASSWORD")
	mailHost     = os.Getenv("MAILTRAP_SMTP_HOST")
	mailPort     = os.Getenv("MAILTRAP_SMTP_PORT")
)

func setupTests() (*AuthHandler, func(), error) {

	ctx := context.Background()

	dbName := "test_db"
	dbUser := "test_user"
	dbPassword := "test_password"

	// postgres test container
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:14-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}

	host, err := pgContainer.Host(ctx)
	if err != nil {
		return &AuthHandler{}, func() {}, err
	}

	port, err := pgContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return &AuthHandler{}, func() {}, err
	}

	dbConfig, err := config.NewDBConfig(host, port.Port(), dbUser, dbPassword, dbName)
	if err != nil {
		return &AuthHandler{}, func() {}, err
	}

	mailClient := config.NewMailConfig(mailFrom, mailPassword, mailHost, mailPort)

	// function to clean up the container
	TerminateContainers := func() {
		// close pgContainer
		if err := pgContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}

	// jwt test config
	jwtConfig := config.NewJWTConfig(secretKey)

	// google client test config
	googleClient := config.NewGoogleAuthClient(googleClientId)

	return NewAuthHandler(dbConfig.DB, googleClient, mailClient, jwtConfig), TerminateContainers, nil
}

func TestSignInSuccess(t *testing.T) {
	handler, TerminateContainers, err := setupTests()
	defer TerminateContainers()

	if err != nil {
		t.Fatal("TestSignInSuccess Failed with Error: ", err.Error())
	}

	name := "Test"
	email := "to@test.com"
	requestBody := map[string]string{"Name": name, "Email": email}
	jsonValue, _ := json.Marshal(requestBody)

	request, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(jsonValue))
	recorder := httptest.NewRecorder()

	signup := http.HandlerFunc(handler.SignIn)
	signup.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

}
