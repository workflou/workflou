package inmem

import (
	"context"
	"workflou/pkg/store"
	"workflou/pkg/workflou"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	Users []*workflou.User
}

func NewUserStore() *UserStore {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	return &UserStore{
		Users: []*workflou.User{
			{
				ID:           "test",
				Name:         "Test",
				Email:        "test@example.com",
				PasswordHash: string(passwordHash),
			},
		},
	}
}

func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (*workflou.User, error) {
	for _, u := range s.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrNotFound
}

func (s *UserStore) SaveUser(ctx context.Context, user *workflou.User) error {
	user.ID = uuid.NewString()
	s.Users = append(s.Users, user)
	return nil
}
