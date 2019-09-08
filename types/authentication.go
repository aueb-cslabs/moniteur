package types

import "github.com/dgrijalva/jwt-go"

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type AuthTokenClaim struct {
	*jwt.StandardClaims
	Username string
}

type ErrorMsg struct {
	Message string `json:"message"`
}
