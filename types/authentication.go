package types

import (
	"github.com/dgrijalva/jwt-go"
)

// AuthToken Bearer auth token
type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthTokenClaim a auth token claim from a user
type AuthTokenClaim struct {
	*jwt.StandardClaims
	Username string
}

// ErrorMsg has an error message
type ErrorMsg struct {
	Message string `json:"message"`
}
