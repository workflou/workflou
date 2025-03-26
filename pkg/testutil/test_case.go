package testutil

import (
	"net/http"
	"net/http/httptest"
	"workflou/pkg/mux"
	"workflou/pkg/store"
	"workflou/pkg/store/inmem"
)

type TestCase struct {
	Store  store.Store
	Server *httptest.Server
	Client *http.Client
}

func NewTestCase() *TestCase {
	store := inmem.New()

	return &TestCase{
		Store:  store,
		Server: httptest.NewServer(mux.New(store)),
	}
}

func (tc *TestCase) Close() {
	tc.Server.Close()
}
