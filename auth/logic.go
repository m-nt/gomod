package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// UserVerifier is what Auth needs from the outside world to process a login
type UserVerifier interface {
	VerifyCredentials(ctx context.Context, email, password string) (int, error)
}

type Service struct {
	secret string
}

func NewService(secret string) *Service {
	return &Service{secret: secret}
}

func (s *Service) GenerateToken(uid int) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	return t.SignedString([]byte(s.secret))
}

func (s *Service) ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	return int(claims["uid"].(float64)), nil
}
