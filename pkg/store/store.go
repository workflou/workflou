package store

import (
	"context"
	"errors"
	"workflou/pkg/workflou"
)

type Store interface {
	GetUserByEmail(ctx context.Context, email string) (*workflou.User, error)
	SaveUser(ctx context.Context, user *workflou.User) error
	GetSessionByID(ctx context.Context, id string) (*workflou.Session, error)
	SaveSession(ctx context.Context, session *workflou.Session) error
	DeleteSession(ctx context.Context, id string) error
}

var (
	ErrNotFound = errors.New("not found")
)
