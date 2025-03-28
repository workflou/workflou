package testutil

import (
	"net/http"
	"time"
	"workflou/pkg/workflou"
)

func CreateSessionAndCookieForUser(user *workflou.User) (*workflou.Session, http.Cookie) {
	session := &workflou.Session{
		ID:        "sessionID",
		User:      user,
		CreatedAt: time.Now(),
	}

	cookie := http.Cookie{
		Name:     string(workflou.SessionKey),
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
	}

	return session, cookie
}
