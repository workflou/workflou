package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"workflou/pkg/workflou"
)

func AssignUserIdFromCookie(sessionStore workflou.SessionStore) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(workflou.SessionKey)
			if err != nil {
				slog.Info("session cookie not found")
				next.ServeHTTP(w, r)
				return
			}

			session, err := sessionStore.GetByID(r.Context(), cookie.Value)
			if err != nil {
				slog.Info("session not found", "sessionID", cookie.Value)
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), workflou.UserKey, session.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func EnsureAuthenticated() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := r.Context().Value(workflou.UserKey).(string)
			if !ok || userID == "" {
				http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
