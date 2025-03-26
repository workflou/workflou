package handler

import (
	"net/http"
	"workflou/pkg/store"
	"workflou/pkg/view"

	"github.com/go-chi/chi/v5"
)

type HomeHandler struct {
	Store store.Store
}

func NewHomeHandler(store store.Store) *HomeHandler {
	return &HomeHandler{
		Store: store,
	}
}

func (h *HomeHandler) Register(mux chi.Router) {
	mux.Get("/", h.HomePage)
}

func (h *HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	view.HomePage().Render(r.Context(), w)
}
