package handler

import (
	"net/http"
	"workflou/pkg/workflou"
)

type AuthHandler struct {
	Users workflou.UserStore
}

func NewAuthHandler(us workflou.UserStore) *AuthHandler {
	return &AuthHandler{Users: us}
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}
