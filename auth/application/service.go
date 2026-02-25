package application

import (
	"context"

	"github.com/m-nt/gomod/auth/domain"
)

type UserProvider interface {
	Exists(ctx context.Context, id int) (bool, error)
}

type TokenProvider interface {
	Generate(domain.Claims) (string, error)
	Parse(token string) (*domain.Claims, error)
}

type Service struct {
	users UserProvider
	jwt   TokenProvider
}

func New(users UserProvider, jwt TokenProvider) *Service {
	return &Service{
		users: users,
		jwt:   jwt,
	}
}

func (s *Service) Login(ctx context.Context, userID int, email string) (string, error) {

	exists, err := s.users.Exists(ctx, userID)
	if err != nil || !exists {
		return "", err
	}

	return s.jwt.Generate(domain.Claims{
		UserID: userID,
		Email:  email,
	})
}

func (s *Service) Validate(token string) (*domain.Claims, error) {
	return s.jwt.Parse(token)
}
