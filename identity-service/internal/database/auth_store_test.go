package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"identity-service/internal/config"
	"identity-service/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// function to create postgres test container and create db config with the created container
func SetupTest() (*AuthStore, func(), error) {
	ctx := context.Background()

	dbName := "test_db"
	dbUser := "test_user"
	dbPassword := "test_password"

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
		return &AuthStore{}, func() {}, err
	}

	port, err := pgContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return &AuthStore{}, func() {}, err
	}

	dbConfig, err := config.NewDBConfig(host, port.Port(), dbUser, dbPassword, dbName)
	if err != nil {
		return &AuthStore{}, func() {}, err
	}

	// function to clean up the container
	TerminateContainers := func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}

	return NewAuthStore(dbConfig.DB), TerminateContainers, nil
}

func TestFindUserByEmail(t *testing.T) {
	store, TerminateContainers, err := SetupTest()
	defer TerminateContainers()

	if err != nil {
		t.Fatal("TestFindUserByEmail Failed with Error: ", err.Error())
	}

	email := "testuser@test.com"
	_, err = store.db.Exec(`INSERT INTO users (email, is_signup_complete, user_created_at) VALUES ($1, $2, $3)`, email, true, time.Now())
	if err != nil {
		t.Log(err)
	}

	user, err := store.FindUserByEmail(email)

	assert.Nil(t, err)
	assert.Equal(t, email, user.Email)
}

func TestCreateUser(t *testing.T) {
	store, TerminateContainers, err := SetupTest()
	defer TerminateContainers()

	if err != nil {
		t.Fatal("TestCreateUser Failed with Error: ", err.Error())
	}

	email := "testuser@test.com"

	user := &models.User{
		Email:            email,
		IsSignupComplete: false,
	}

	_, err = store.CreateUser(user)
	assert.Nil(t, err)

	err = store.db.QueryRow(`SELECT * FROM users WHERE email = $1`, email).
		Scan(&user.Id, &user.Name, &user.Email, &user.IsSignupComplete, &user.CreatedAt)

	assert.Nil(t, err)
	assert.Equal(t, "", user.Name)
	assert.Equal(t, email, user.Email)
}

func TestUpdateUser(t *testing.T) {
	store, TerminateContainers, err := SetupTest()
	defer TerminateContainers()

	if err != nil {
		t.Fatal("TestUpdateUser Failed with Error: ", err.Error())
	}

	name := "Test User"
	email := "testuser@test.com"

	_, err = store.db.Exec(`INSERT INTO users (email, is_signup_complete, user_created_at) VALUES ($1, $2, $3)`, email, true, time.Now())
	if err != nil {
		t.Fatal("TestUpdateUser Failed with Error: ", err.Error())
	}

	user := &models.User{
		Name:             name,
		Email:            email,
		IsSignupComplete: false,
	}
	err = store.UpdateUser(user)
	assert.Nil(t, err)

	err = store.db.QueryRow(`SELECT * FROM users WHERE email = $1`, email).
		Scan(&user.Id, &user.Name, &user.Email, &user.IsSignupComplete, &user.CreatedAt)

	assert.Nil(t, err)
	assert.Equal(t, name, user.Name)
}

func TestFindOtpByEmail(t *testing.T) {
	store, TerminateContainers, err := SetupTest()
	defer TerminateContainers()

	if err != nil {
		t.Fatal("TestFindOtpByEmail Failed with Error: ", err.Error())
	}

	email := "testuser@test.com"
	otpValue := "123456"
	_, err = store.db.Exec(`INSERT INTO otp (email, otp, expiration) VALUES ($1, $2, $3)`, email, otpValue, time.Now())
	if err != nil {
		t.Fatal("TestFindOtpByEmail Failed with Error: ", err.Error())
	}

	otp, err := store.FindOtpByEmail(email)

	assert.Nil(t, err)
	assert.Equal(t, email, otp.Email)
	assert.Equal(t, otpValue, otp.Value)
}

func TestCreateOtp(t *testing.T) {
	store, TerminateContainers, err := SetupTest()
	defer TerminateContainers()

	if err != nil {
		t.Fatal("TestCreateOtp Failed with Error: ", err.Error())
	}

	email := "testuser@test.com"
	otpValue := "123456"

	otp := &models.OTP{
		Email:      email,
		Value:      otpValue,
		Expiration: time.Now(),
	}

	err = store.CreateOtp(otp)
	assert.Nil(t, err)

	err = store.db.QueryRow(`SELECT * FROM otp WHERE email = $1`, email).
		Scan(&otp.Id, &otp.Email, &otp.Value, &otp.Expiration)

	assert.Nil(t, err)
	assert.Equal(t, email, otp.Email)
}

func TestDeleteOtp(t *testing.T) {
	store, TerminateContainers, err := SetupTest()
	defer TerminateContainers()

	if err != nil {
		t.Fatal("TestDeleteOtp Failed with Error: ", err.Error())
	}

	email := "testuser@test.com"
	otpValue := "123456"
	_, err = store.db.Exec(`INSERT INTO otp (email, otp, expiration) VALUES ($1, $2, $3)`, email, otpValue, time.Now())
	if err != nil {
		t.Fatal("TestDeleteOtp Failed with Error: ", err.Error())
	}

	err = store.DeleteOtp(email)
	assert.Nil(t, err)

	var otp models.OTP
	err = store.db.QueryRow(`SELECT * FROM otp WHERE email = $1`, email).
		Scan(&otp.Id, &otp.Email, &otp.Value, &otp.Expiration)

	assert.Equal(t, sql.ErrNoRows, err)
}
