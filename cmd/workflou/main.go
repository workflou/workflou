package main

import (
	"log/slog"
	"workflou/pkg/server"
)

func main() {
	s := server.New(":4000")
	slog.Info("http://localhost:4000")
	s.ListenAndServe()
}
