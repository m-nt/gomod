package application

import (
	"context"

	"github.com/m-nt/gomod/users/domain"
	"github.com/m-nt/gomod/users/infrastructure/ent"
)

type Service struct {
	client *ent.Client
}

func New(client *ent.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Create(ctx context.Context, email, name string) (*domain.User, error) {
	u, err := domain.NewUser(email, name)
	if err != nil {
		return nil, err
	}

	e, err := s.client.User.
		Create().
		SetEmail(u.Email).
		SetName(u.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	u.ID = e.ID
	return u, nil
}

func (s *Service) Get(ctx context.Context, id int) (*domain.User, error) {
	e, err := s.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
	}, nil
}
