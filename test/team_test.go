package test

import (
	"net/http"
	"testing"
	"workflou/pkg/store/inmem"
	"workflou/pkg/testutil"
	"workflou/pkg/workflou"
)

func TestTeamIsRequired(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	store := tc.Store.(inmem.Store)
	user := &workflou.User{
		ID:           "teamless-user",
		Name:         "Teamless User",
		Email:        "teamless@example.com",
		PasswordHash: "passwordHash",
	}
	store.Users = append(store.Users, user)
	session, cookie := testutil.CreateSessionAndCookieForUser(user)
	store.Sessions = append(store.Sessions, session)

	req, _ := http.NewRequest("GET", tc.Server.URL+"/", nil)
	req.AddCookie(&cookie)

	res, err := tc.Server.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.Request.Response == nil {
		t.Fatalf("Expected response to be set")
	}

	if statusCode := res.Request.Response.StatusCode; statusCode != http.StatusSeeOther {
		t.Fatalf("Expected status code %d, got %d", http.StatusSeeOther, statusCode)
	}

	if location := res.Request.Response.Header.Get("Location"); location != "/teams/new" {
		t.Fatalf("Expected redirect location %s, got %s", "/teams/new", location)
	}
}
