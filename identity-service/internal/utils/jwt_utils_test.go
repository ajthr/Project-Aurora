package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractBearerToken(t *testing.T) {
	// Test Extracting Bearer Token from Authorization Header String
	assert.NotEmpty(t, ExtractBearerToken("Bearer jwt_token"))
	assert.Empty(t, ExtractBearerToken("Bearer "))
	assert.Empty(t, ExtractBearerToken("jwt_token"))
	assert.Empty(t, ExtractBearerToken(""))
}
