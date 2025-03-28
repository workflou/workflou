package mux

import (
	"net/http"
	"time"
	"workflou/assets/css"
	"workflou/assets/img"
	"workflou/assets/js"
	"workflou/pkg/handler"
	"workflou/pkg/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(store store.Store) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Timeout(time.Second * 5))

	r.Handle("/css/*", http.StripPrefix("/css/", http.FileServerFS(css.FS)))
	r.Handle("/js/*", http.StripPrefix("/js/", http.FileServerFS(js.FS)))
	r.Handle("/img/*", http.StripPrefix("/img/", http.FileServerFS(img.FS)))

	r.Group(func(r chi.Router) {
		r.Use(AssignUserFromCookie(store))
		r.Use(EnsureAuthenticated())
		r.Use(AssignTeams(store))
		r.Use(EnsureHasTeam())

		handler.NewHomeHandler(store).Register(r)
		handler.NewLogoutHandler(store).Register(r)
	})

	r.Group(func(r chi.Router) {
		handler.NewLoginHandler(store).Register(r)
	})

	return r
}
