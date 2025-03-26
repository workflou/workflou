package inmem

import (
	"context"
	"workflou/pkg/store"
	"workflou/pkg/workflou"

	"github.com/google/uuid"
)

type SessionStore struct {
	Sessions []*workflou.Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		Sessions: []*workflou.Session{},
	}
}

func (s *SessionStore) SaveSession(ctx context.Context, session *workflou.Session) error {
	session.ID = uuid.NewString()
	s.Sessions = append(s.Sessions, session)
	return nil
}

func (s *SessionStore) GetSessionByID(ctx context.Context, id string) (*workflou.Session, error) {
	for _, session := range s.Sessions {
		if session.ID == id {
			return session, nil
		}
	}
	return nil, store.ErrNotFound
}

func (s *SessionStore) DeleteSession(ctx context.Context, id string) error {
	for i, session := range s.Sessions {
		if session.ID == id {
			s.Sessions = append(s.Sessions[:i], s.Sessions[i+1:]...)
			return nil
		}
	}
	return store.ErrNotFound
}
