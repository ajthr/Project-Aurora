package database

import (
	"database/sql"
	"identity-service/internal/models"
	"time"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(conn *sql.DB) *AuthStore {
	return &AuthStore{
		db: conn,
	}
}

func (s *AuthStore) FindUserById(userId int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE user_id = $1`
	row := s.db.QueryRow(query, userId)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.IsSignupComplete, &user.CreatedAt); err != nil {
		return &user, err
	}
	return &user, nil
}

func (s *AuthStore) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE email = $1"
	row := s.db.QueryRow(query, email)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.IsSignupComplete, &user.CreatedAt); err != nil {
		return &user, err
	}
	return &user, nil
}

func (s *AuthStore) CreateUser(user *models.User) error {
	query := `INSERT INTO users (name, email, is_signup_complete, user_created_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, user.Name, user.Email, false, time.Now())
	return err
}

func (s *AuthStore) UpdateSignUpStatus(email string) error {
	query := `UPDATE users SET is_signup_complete = true WHERE email = $1`
	_, err := s.db.Exec(query, email)
	return err
}

func (s *AuthStore) FindOtpByEmail(email string) (*models.OTP, error) {
	var otp models.OTP
	query := `SELECT * FROM otp WHERE email = $1`
	row := s.db.QueryRow(query, email)
	if err := row.Scan(&otp.Id, &otp.Email, &otp.Value, &otp.Expiration); err != nil {
		return &otp, err
	}
	return &otp, nil
}

func (s *AuthStore) DeleteOtp(email string) error {
	query := `DELETE FROM otp WHERE email = $1`
	_, err := s.db.Exec(query, email)
	return err
}

func (s *AuthStore) CreateOtp(otp *models.OTP) error {
	query := `INSERT INTO otp (email, otp, expiration) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, otp.Email, otp.Value, otp.Expiration)
	return err
}
