package jwt

import (
	"errors"
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
func (t *tokenService) GenerateToken(username, role string, expire time.Time) (string, error) {
	// Set custom claims
	claims := &entity.JwtCustomClaims{
		Username: username,
		Role:     role,
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

// ValidateToken will validate a token
func (t *tokenService) ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(t.secret), nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return false
	}

	return true
}
