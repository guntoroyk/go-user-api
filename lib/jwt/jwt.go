package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/guntoroyk/go-user-api/entity"
)

type tokenService struct {
	secret string
}

// NewTokenService will create new an TokenService object representation of TokenServiceItf interface
func NewTokenService(secret string) TokenServiceItf {
	return &tokenService{
		secret: secret,
	}
}

// GenerateToken will generate a token
func (t *tokenService) GenerateToken(username, role string, tokenType entity.TokenType, expire time.Time) (string, error) {
	// Set custom claims
	claims := &entity.JwtCustomClaims{
		Username: username,
		Role:     role,
		Type:     tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expire,
			},
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
