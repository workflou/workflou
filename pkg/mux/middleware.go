package mux

import (
	"context"
	"net/http"
	"workflou/pkg/store"
	"workflou/pkg/workflou"
)

type Middleware func(http.Handler) http.Handler

func AssignUserFromCookie(store store.Store) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(string(workflou.SessionKey))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			session, err := store.GetSessionByID(r.Context(), cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), workflou.UserKey, session.User)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func EnsureAuthenticated() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(workflou.UserKey).(*workflou.User)
			if !ok || user == nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
