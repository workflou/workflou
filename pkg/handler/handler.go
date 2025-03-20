package handler

import (
	"net/http"
	"workflou/pkg/inmem"
)

func New() *http.ServeMux {
	us := inmem.NewUserStore()

	mux := http.NewServeMux()

	ah := NewAuthHandler(us)
	mux.HandleFunc("/login", ah.LoginPage())

	return mux
}
