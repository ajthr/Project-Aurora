package utils

import "strings"

func ExtractBearerToken(authorizationHeader string) string {
	if authorizationHeader != "" && strings.HasPrefix(authorizationHeader, "Bearer ") {
		return strings.TrimPrefix(authorizationHeader, "Bearer ")
	}
	return ""
}
