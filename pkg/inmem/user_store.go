package inmem

import (
	"context"
	"workflou/pkg/workflou"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	users []*workflou.User
}

func NewUserStore() *UserStore {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	return &UserStore{
		users: []*workflou.User{
			{
				ID:           "test",
				Name:         "Test",
				Email:        "test@example.com",
				PasswordHash: string(passwordHash),
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
