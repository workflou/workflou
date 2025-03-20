package inmem

import (
	"context"
	"workflou/pkg/workflou"

	"github.com/google/uuid"
)

type UserStore struct {
	users []*workflou.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: []*workflou.User{
			{
				ID:    "alice",
				Name:  "Alice",
				Email: "alice@example.com",
			},
		},
	}
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*workflou.User, error) {
	for _, u := range s.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, workflou.ErrUserNotFound
}

func (s *UserStore) Save(ctx context.Context, user *workflou.User) error {
	user.ID = uuid.NewString()
	s.users = append(s.users, user)
	return nil
}
