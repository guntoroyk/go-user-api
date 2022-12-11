package jwt

import (
	"time"

	"github.com/guntoroyk/go-user-api/entity"
)

// TokenServiceItf is an interface for token service
type TokenServiceItf interface {
	GenerateToken(username, role string, tokenType entity.TokenType, expire time.Time) (string, error)
}
