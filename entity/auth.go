package entity

import "github.com/golang-jwt/jwt/v4"

// Token is the struct for a token
type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// JwtCustomClaims is the struct for a custom claims
type JwtCustomClaims struct {
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Type     TokenType `json:"type"`
	jwt.RegisteredClaims
}

// TokenType is the type for token type
type TokenType string

// Token types
const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)
