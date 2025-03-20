package auth

import (
	"context"
	"net/http"
	"workflou/pkg/workflou"
)

type Middleware func(http.Handler) http.Handler

func AssignUserIdFromCookie(sessionStore workflou.SessionStore) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			session, err := sessionStore.GetByID(r.Context(), cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", session.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func EnsureAuthenticated() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := r.Context().Value("user_id").(string)
			if !ok || userID == "" {
				http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
