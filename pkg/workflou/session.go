package workflou

import (
	"time"
)

type Session struct {
	ID        string
	User      *User
	CreatedAt time.Time
}
