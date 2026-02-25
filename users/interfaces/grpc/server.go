package grpc

import (
	"context"

	"github.com/m-nt/gomod/users/application"
)

type Server struct {
	svc *application.Service
	UnimplementedUsersServiceServer
}

func New(svc *application.Service) *Server {
	return &Server{svc: svc}
}

func (s *Server) Create(ctx context.Context, r *CreateUserRequest) (*UserResponse, error) {
	u, err := s.svc.Create(ctx, r.Email, r.Name)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		Id:    int32(u.ID),
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
