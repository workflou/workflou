package viewmodel

import "net/mail"

type LoginForm struct {
	Email    string
	Password string

	Errors map[string]string
}

func (f *LoginForm) Valid() bool {
	if f.Email == "" {
		f.Errors["Email"] = "Email is required"
	} else if _, err := mail.ParseAddress(f.Email); err != nil {
		f.Errors["Email"] = "Email is invalid"
	}

	if f.Password == "" {
		f.Errors["Password"] = "Password is required"
	}

	return len(f.Errors) == 0
}
