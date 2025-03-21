package server

import (
	"net/http"
	"workflou/assets/css"
	"workflou/assets/js"
	"workflou/pkg/handler"
	"workflou/pkg/inmem"
	"workflou/pkg/middleware"
)

func New(addr string) *http.Server {
	userStore := inmem.NewUserStore()
	sessionStore := inmem.NewSessionStore()

	commonMiddleware := middleware.NewStack(
		middleware.RequestID(),
		middleware.LogRequest(),
		middleware.Recover(),
	)

	authMiddleware := middleware.NewStack(
		middleware.AssignUserIdFromCookie(sessionStore),
		middleware.EnsureAuthenticated(),
	)

	guestMiddleware := middleware.NewStack()

	mux := http.NewServeMux()
	authMux := http.NewServeMux()

	authHandler := handler.NewAuthHandler(userStore, sessionStore)
	authMux.HandleFunc("GET /login", authHandler.LoginPage())
	authMux.HandleFunc("POST /login", authHandler.LoginForm())

	guestMux := http.NewServeMux()
	guestMux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})

	mux.Handle("/", authMiddleware(guestMux))
	mux.Handle("/auth/", http.StripPrefix("/auth", guestMiddleware(authMux)))

	mux.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oh no")
	})

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServerFS(css.FS)))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServerFS(js.FS)))

	return &http.Server{
		Addr:    addr,
		Handler: commonMiddleware(mux),
	}
}
