package workflou

import (
	"context"
	"errors"
)

type User struct {
	ID    string
	Name  string
	Email string
}

type UserStore interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

var ErrUserNotFound = errors.New("user not found")
