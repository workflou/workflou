package test

import (
	"net/http"
	"testing"
	"time"
	"workflou/pkg/store/inmem"
	"workflou/pkg/testutil"
	"workflou/pkg/workflou"
)

func TestLogout_ByGuest(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	r, _ := tc.Server.Client().Get(tc.Server.URL + "/logout")

	if r.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, r.StatusCode)
	}

	if r.Request == nil || r.Request.Response == nil {
		t.Fatal("Expected response to be set")
	}

	if statusCode := r.Request.Response.StatusCode; statusCode != http.StatusSeeOther {
		t.Errorf("Expected status code %d, got %d", http.StatusSeeOther, statusCode)
	}

	if location := r.Request.Response.Header.Get("Location"); location != "/login" {
		t.Errorf("Expected redirect location %s, got %s", "/login", location)
	}
}

func TestLogout_ByAuthenticatedUser(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	store := tc.Store.(inmem.Store)
	team := &workflou.Team{
		ID:   "team-1",
		Name: "Team #1",
	}
	user := &workflou.User{
		ID:           "1",
		Email:        "test@example.com",
		PasswordHash: "passwordHash",
		Teams:        []*workflou.Team{team},
		CurrentTeam:  team,
	}
	store.Users = append(store.Users, user)
	session, cookie := testutil.CreateSessionAndCookieForUser(user)
	store.Sessions = append(store.Sessions, session)

	req, _ := http.NewRequest("GET", tc.Server.URL+"/logout", nil)
	req.AddCookie(&cookie)

	r, err := tc.Server.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if r.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, r.StatusCode)
	}

	if r.Request == nil || r.Request.Response == nil {
		t.Fatal("Expected response to be set")
	}

	if statusCode := r.Request.Response.StatusCode; statusCode != http.StatusSeeOther {
		t.Errorf("Expected status code %d, got %d", http.StatusSeeOther, statusCode)
	}

	if location := r.Request.Response.Header.Get("Location"); location != "/login" {
		t.Errorf("Expected redirect location %s, got %s", "/login", location)
	}

	if len(store.Sessions) != 0 {
		t.Errorf("Expected all sessions to be deleted. Got %v", store.Sessions[0])
	}
}

func TestLogout_OtherSessionsAreNotAffected(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	store := tc.Store.(inmem.Store)

	team := &workflou.Team{
		ID:   "team-1",
		Name: "Team #1",
	}
	user1 := &workflou.User{
		ID:           "1",
		Email:        "test1@example.com",
		PasswordHash: "passwordHash",
		Teams:        []*workflou.Team{team},
		CurrentTeam:  team,
	}
	user2 := &workflou.User{
		ID:           "2",
		Email:        "test2@example.com",
		PasswordHash: "passwordHash",
		Teams:        []*workflou.Team{team},
		CurrentTeam:  team,
	}
	store.Users = append(store.Users, user1, user2)

	session1 := &workflou.Session{
		ID:        "sessionID1",
		User:      user1,
		CreatedAt: time.Now(),
	}
	session2 := &workflou.Session{
		ID:        "sessionID2",
		User:      user2,
		CreatedAt: time.Now(),
	}
	store.Sessions = append(store.Sessions, session1, session2)

	cookie := http.Cookie{
		Name:     string(workflou.SessionKey),
		Value:    session1.ID,
		Path:     "/",
		HttpOnly: true,
	}

	req, _ := http.NewRequest("GET", tc.Server.URL+"/logout", nil)
	req.AddCookie(&cookie)

	_, err := tc.Server.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if len(store.Sessions) != 1 {
		t.Errorf("Expected only one session to be deleted. Got %v", store.Sessions)
	}

	if store.Sessions[0].ID != session2.ID {
		t.Errorf("Expected session %v to be deleted. Got %v", session2, store.Sessions[0])
	}
}
