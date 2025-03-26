package main

import (
	"fmt"
	"net/http"
	"workflou/pkg/mux"
	"workflou/pkg/store/inmem"
)

func main() {
	store := inmem.New()
	r := mux.New(store)

	fmt.Println("http://localhost:4000")
	http.ListenAndServe(":4000", r)
}
