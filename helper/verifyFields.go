package helper

import (
	"errors"
)

// VerifyFields checks if required string fields in a struct-like object are non-empty
func VerifyFields(email, username, name, phoneNumber, password string) error {
	if email == "" {
		return errors.New("email is required")
	}

	if username == "" {
		return errors.New("username is required")
	}

	if name == "" {
		return errors.New("name is required")
	}

	if phoneNumber == "" {
		return errors.New("phone number is required")
	}

	if password == "" {
		return errors.New("password is required")
	}

	return nil
}
