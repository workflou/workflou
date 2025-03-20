package inmem

import (
	"context"
	"workflou/pkg/workflou"

	"github.com/google/uuid"
)

type SessionStore struct {
	sessions []*workflou.Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: []*workflou.Session{},
	}
}

func (s *SessionStore) Save(ctx context.Context, session *workflou.Session) error {
	session.ID = uuid.NewString()
	s.sessions = append(s.sessions, session)
	return nil
}

func (s *SessionStore) GetByID(ctx context.Context, id string) (*workflou.Session, error) {
	for _, session := range s.sessions {
		if session.ID == id {
			return session, nil
		}
	}
	return nil, workflou.ErrSessionNotFound
}
