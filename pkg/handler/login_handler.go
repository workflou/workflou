package handler

import (
	"net/http"
	"time"
	"workflou/pkg/store"
	"workflou/pkg/view"
	"workflou/pkg/viewmodel"
	"workflou/pkg/workflou"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Store store.Store
}

func NewLoginHandler(store store.Store) *LoginHandler {
	return &LoginHandler{
		Store: store,
	}
}

func (h *LoginHandler) Register(mux chi.Router) {
	mux.Get("/login", h.LoginPage)
	mux.Post("/login", h.LoginForm)
}

func (h *LoginHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	view.LoginPage(&viewmodel.LoginForm{}).Render(r.Context(), w)
}

func (h *LoginHandler) LoginForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	f := &viewmodel.LoginForm{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Errors:   make(map[string]string),
	}

	if !f.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		view.LoginPage(f).Render(r.Context(), w)
		return
	}

	user, err := h.Store.GetUserByEmail(r.Context(), f.Email)
	if err != nil || user == nil {
		f.Errors["Email"] = "Invalid email or password"
		w.WriteHeader(http.StatusBadRequest)
		view.LoginPage(f).Render(r.Context(), w)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(f.Password)) != nil {
		f.Errors["Email"] = "Invalid email or password"
		w.WriteHeader(http.StatusBadRequest)
		view.LoginPage(f).Render(r.Context(), w)
		return
	}

	s := &workflou.Session{User: user, CreatedAt: time.Now()}
	h.Store.SaveSession(r.Context(), s)

	http.SetCookie(w, &http.Cookie{
		Path:     "/",
		Name:     string(workflou.SessionKey),
		Value:    s.ID,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // todo: production
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
