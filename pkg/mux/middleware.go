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

func AssignCurrentTeam(store store.Store) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(workflou.UserKey).(*workflou.User)
			if !ok || user == nil {
				next.ServeHTTP(w, r)
				return
			}

			if user.CurrentTeam != nil {
				ctx := context.WithValue(r.Context(), workflou.TeamKey, user.CurrentTeam)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			if len(user.Teams) == 0 {
				next.ServeHTTP(w, r)
				return
			}

			user.CurrentTeam = user.Teams[0]
			ctx := context.WithValue(r.Context(), workflou.TeamKey, user.CurrentTeam)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func EnsureHasTeam() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			team, ok := r.Context().Value(workflou.TeamKey).(*workflou.Team)
			if !ok || team == nil {
				http.Redirect(w, r, "/teams/new", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
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
