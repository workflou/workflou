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
	teams := []*workflou.Team{
		&workflou.Team{
			ID:   "team-1",
			Name: "Team #1",
		},
		&workflou.Team{
			ID:   "team-2",
			Name: "This is a very long team name",
		},
		&workflou.Team{
			ID:   "team-3",
			Name: "Team #3",
		},
	}

	return &UserStore{
		Users: []*workflou.User{
			{
				ID:           "test",
				Name:         "This is a very long user name",
				Email:        "test@example.com",
				PasswordHash: string(passwordHash),
				Teams:        teams,
				CurrentTeam:  teams[1],
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
