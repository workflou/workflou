package workflou

import (
	"context"
	"errors"
	"time"
)

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
}

type SessionStore interface {
	GetByID(ctx context.Context, id string) (*Session, error)
	Save(ctx context.Context, session *Session) error
}

var ErrSessionNotFound = errors.New("session not found")
