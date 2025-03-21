package handler

import (
	"net/http"
	"workflou/pkg/inmem"
	"workflou/pkg/middleware"
)

func New() *http.ServeMux {
	userStore := inmem.NewUserStore()
	sessionStore := inmem.NewSessionStore()

	commonMiddleware := middleware.NewStack(
		middleware.LogRequest(),
	)

	authMiddleware := middleware.NewStack(
		commonMiddleware,
		middleware.AssignUserIdFromCookie(sessionStore),
		middleware.EnsureAuthenticated(),
	)

	guestMiddleware := middleware.NewStack(
		commonMiddleware,
	)

	mux := http.NewServeMux()
	authMux := http.NewServeMux()

	authHandler := NewAuthHandler(userStore, sessionStore)
	authMux.HandleFunc("GET /login", authHandler.LoginPage())
	authMux.HandleFunc("POST /login", authHandler.LoginForm())

	guestMux := http.NewServeMux()
	guestMux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})

	mux.Handle("/", authMiddleware(guestMux))
	mux.Handle("/auth/", http.StripPrefix("/auth", guestMiddleware(authMux)))

	return mux
}
