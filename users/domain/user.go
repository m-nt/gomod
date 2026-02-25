package domain

import "errors"

var ErrInvalidUser = errors.New("invalid user")

type User struct {
	ID    int
	Email string
	Name  string
}

func NewUser(email, name string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidUser
	}
	return &User{Email: email, Name: name}, nil
}
