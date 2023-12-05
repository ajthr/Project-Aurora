package middlewares

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"identity-service/internal/config"
	"identity-service/internal/utils"
)

// jwt authenticator middleware
func Authenticator(config *config.JWTConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := utils.ExtractBearerToken(r.Header.Get("Authorization"))

			if tokenString != "" {
				token, err := config.TokenParser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("Signing Method Error!")
					}
					return config.SecretKey, nil
				})
				if err == nil {
					claims, ok := token.Claims.(jwt.MapClaims)
					if ok && token.Valid {
						userId, ok := claims["user_id"].(string)
						if ok && userId != "" {
							w.Header().Set("X-Authenticated-User-ID", userId)
							return
						}
					}
				}
			}
			w.WriteHeader(http.StatusUnauthorized)
		})
	}
}
