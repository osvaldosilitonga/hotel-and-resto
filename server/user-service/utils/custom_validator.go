package utils

import (
	"net/mail"
)

func LoginValidator(email, password string) (bool, string) {
	if email == "" || password == "" {
		return false, "user or password cannot be null"
	}

	if len(email) < 8 || len(password) < 6 {
		return false, "please insert email min length 8, password min length 6"
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, "please insert correct email"
	}

	return true, ""
}
