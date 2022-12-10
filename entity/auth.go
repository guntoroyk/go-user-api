package entity

import "github.com/golang-jwt/jwt/v4"

// Token is the struct for a token
type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// JwtCustomClaims is the struct for a custom claims
type JwtCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
