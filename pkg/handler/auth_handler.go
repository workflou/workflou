package handler

import (
	"html/template"
	"log/slog"
	"net/http"
	"time"
	"workflou/assets/html"
	"workflou/pkg/form"
	"workflou/pkg/workflou"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Users    workflou.UserStore
	Sessions workflou.SessionStore
}

func NewAuthHandler(us workflou.UserStore, ss workflou.SessionStore) *AuthHandler {
	return &AuthHandler{Users: us, Sessions: ss}
}

func (h *AuthHandler) LoginPage() http.HandlerFunc {
	t := template.Must(template.ParseFS(html.FS, "layout.html", "login.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, form.NewLoginForm(r))
	}
}

func (h *AuthHandler) LoginForm() http.HandlerFunc {
	t := template.Must(template.ParseFS(html.FS, "layout.html", "login.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		f := form.NewLoginForm(r)
		if !f.Valid() {
			t.Execute(w, f)
			return
		}

		u, err := h.Users.GetByEmail(r.Context(), f.Email)
		if err != nil {
			f.Errors["Email"] = "invalid email or password"
			t.Execute(w, f)
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(f.Password)) != nil {
			f.Errors["Email"] = "invalid email or password"
			t.Execute(w, f)
			return
		}

		s := &workflou.Session{UserID: u.ID, CreatedAt: time.Now()}
		h.Sessions.Save(r.Context(), s)

		slog.Info("new session", "sessionID", s.ID)

		http.SetCookie(w, &http.Cookie{
			Path:     "/",
			Name:     workflou.SessionKey,
			Value:    s.ID,
			SameSite: http.SameSiteLaxMode,
			Secure:   false, // todo: production
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
