package server

import (
	"net/http"
	"workflou/pkg/handler"
)

func New(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler.New(),
	}
}
