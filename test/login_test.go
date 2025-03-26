package test

import (
	"net/http"
	"net/url"
	"testing"
	"workflou/pkg/store/inmem"
	"workflou/pkg/testutil"
	"workflou/pkg/workflou"

	"golang.org/x/crypto/bcrypt"
)

func TestLogin_PageRedirect(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	r, _ := tc.Server.Client().Get(tc.Server.URL + "/")
	if statusCode := r.Request.Response.StatusCode; statusCode != http.StatusSeeOther {
		t.Errorf("Expected status code %d, got %d", http.StatusSeeOther, statusCode)
	}
	if location := r.Request.Response.Header.Get("Location"); location != "/login" {
		t.Errorf("Expected redirect location %s, got %s", "/login", location)
	}
}

func TestLogin_PageWorks(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	r, _ := tc.Server.Client().Get(tc.Server.URL + "/login")
	if statusCode := r.StatusCode; statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, statusCode)
	}
}

func TestLogin_FormValidationRequiredFields(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	r, _ := tc.Server.Client().Post(tc.Server.URL+"/login", "application/x-www-form-urlencoded", nil)
	if statusCode := r.StatusCode; statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, statusCode)
	}
}

func TestLogin_FormValidationEmailFormatInvalid(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	f := url.Values{
		"email": {"invalid-email"},
	}

	r, _ := tc.Server.Client().PostForm(tc.Server.URL+"/login", f)
	if statusCode := r.StatusCode; statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, statusCode)
	}
}

func TestLogin_UserDoesNotExist(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	f := url.Values{
		"email":    {"invalid@user.com"},
		"password": {"password"},
	}

	r, _ := tc.Server.Client().PostForm(tc.Server.URL+"/login", f)
	if statusCode := r.StatusCode; statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, statusCode)
	}
}

func TestLogin_InvalidPassword(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	store := tc.Store.(inmem.Store)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	store.Users = append(store.Users, &workflou.User{
		ID:           "user-id",
		Email:        "user@example.com",
		PasswordHash: string(passwordHash),
	})

	f := url.Values{
		"email":    {"user@example.com"},
		"password": {"invalid-password"},
	}

	r, _ := tc.Server.Client().PostForm(tc.Server.URL+"/login", f)
	if statusCode := r.StatusCode; statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, statusCode)
	}

	if len(store.Sessions) != 0 {
		t.Error("Expected session to not be saved")
	}
}

func TestLogin_SuccessfulLogin(t *testing.T) {
	tc := testutil.NewTestCase()
	defer tc.Close()

	store := tc.Store.(inmem.Store)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	store.Users = append(store.Users, &workflou.User{
		ID:           "user-id",
		Email:        "user@example.com",
		PasswordHash: string(passwordHash),
	})

	f := url.Values{
		"email":    {"user@example.com"},
		"password": {"password"},
	}

	r, _ := tc.Server.Client().PostForm(tc.Server.URL+"/login", f)

	if r.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, r.StatusCode)
	}

	if r.Request.Response.StatusCode != http.StatusSeeOther {
		t.Fatalf("Expected status code %d, got %d", http.StatusSeeOther, r.Request.Response.StatusCode)
	}

	// if location := r.Request.Response.Header.Get("Location"); location != "/" {
	// 	t.Errorf("Expected redirect location %s, got %s", "/", location)
	// }

	if len(store.Sessions) != 1 {
		t.Error("Expected session to be saved")
	}
}
