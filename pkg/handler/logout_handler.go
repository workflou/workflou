package handler

import (
	"net/http"
	"workflou/pkg/store"
	"workflou/pkg/workflou"

	"github.com/go-chi/chi/v5"
)

type LogoutHandler struct {
	Store store.Store
}

func NewLogoutHandler(store store.Store) *LogoutHandler {
	return &LogoutHandler{Store: store}
}

func (h *LogoutHandler) Register(mux chi.Router) {
	mux.HandleFunc("/logout", h.Logout)
}

func (h *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie(string(workflou.SessionKey))
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	h.Store.DeleteSession(r.Context(), sessionCookie.Value)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
