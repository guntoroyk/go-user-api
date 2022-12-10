package jwt

import "time"

// TokenServiceItf is an interface for token service
type TokenServiceItf interface {
	GenerateToken(username, role string, expire time.Time) (string, error)
	ValidateToken(tokenString string) bool
}
