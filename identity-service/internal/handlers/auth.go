package handlers

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"identity-service/internal/config"
	"identity-service/internal/database"
	"identity-service/internal/models"
)

type AuthHandler struct {
	store        *database.AuthStore
	googleClient *config.GoogleAuthClient
	mailClient   *config.MailConfig
	jwtConfig    *config.JWTConfig
}

func NewAuthHandler(conn *sql.DB, googleClient *config.GoogleAuthClient, mailClient *config.MailConfig, jwtConfig *config.JWTConfig) *AuthHandler {
	authStore := database.NewAuthStore(conn)
	return &AuthHandler{
		store:        authStore,
		googleClient: googleClient,
		mailClient:   mailClient,
		jwtConfig:    jwtConfig,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var request models.SignInRequest
	var err error

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = h.store.FindUserByEmail(request.Email); err != nil {
		if err == sql.ErrNoRows {
			user := &models.User{
				Name:             request.Name,
				Email:            request.Email,
				IsSignupComplete: false,
			}
			if _, err = h.store.CreateUser(user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err = h.store.DeleteOtp(request.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	otpValue := strconv.Itoa(rand.Intn(900000) + 100000)

	otp := &models.OTP{
		Email:      request.Email,
		Value:      otpValue,
		Expiration: time.Now().Add(time.Minute * 5),
	}

	if err = h.store.CreateOtp(otp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mailData := &models.OTPMailData{OTP: otpValue}
	if err = h.mailClient.SendOtpEmail(request.Email, mailData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// function to signin and signup with google auth
func (h *AuthHandler) GoogleSignIn(w http.ResponseWriter, r *http.Request) {

	var request models.GoogleSigninRequest
	var user *models.User
	var token string
	var response []byte
	var err error

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.googleClient.VerifyGoogleToken(request.Token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user, err = h.store.FindUserByEmail(h.googleClient.Claims.Email); err != nil {
		if err == sql.ErrNoRows {
			user := &models.User{
				Name:             h.googleClient.Claims.Name,
				Email:            h.googleClient.Claims.Email,
				IsSignupComplete: true,
			}
			id, err := h.store.CreateUser(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			user.Id = id
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	token, err = h.jwtConfig.CreateToken(user.Id)
	if err == nil {
		response, err = json.Marshal(&models.JWTTokenResponse{
			Token:            token,
			IsSignupComplete: user.IsSignupComplete,
		})
		if err == nil {
			h.store.DeleteOtp(h.googleClient.Claims.Email)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
			return
		}
	}
}

// function to validate otp for signin and signup
func (h *AuthHandler) ValidateOtp(w http.ResponseWriter, r *http.Request) {

	var request models.OTPValidationRequest
	var user *models.User
	var otp *models.OTP
	var token string
	var response []byte
	var err error

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	otp, err = h.store.FindOtpByEmail(request.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if otp.Value == request.OTP && time.Now().Before(otp.Expiration) {
		if user, err = h.store.FindUserByEmail(request.Email); err == nil {
			token, err = h.jwtConfig.CreateToken(user.Id)
			if err == nil {
				response, err = json.Marshal(&models.JWTTokenResponse{
					Token:            token,
					IsSignupComplete: user.IsSignupComplete,
				})
				if err == nil {
					h.store.DeleteOtp(request.Email)
					w.Header().Set("Content-Type", "application/json")
					w.Write(response)
					return
				}
			}
		}
		http.Error(w, "Expired or Invalid Token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
