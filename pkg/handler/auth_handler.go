package handler

import (
	"html/template"
	"net/http"
	"workflou/assets/html"
	"workflou/pkg/workflou"
)

type AuthHandler struct {
	Users workflou.UserStore
}

func NewAuthHandler(us workflou.UserStore) *AuthHandler {
	return &AuthHandler{Users: us}
}

func (h *AuthHandler) LoginPage() http.HandlerFunc {
	t := template.Must(template.ParseFS(html.FS, "layout.html", "login.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
