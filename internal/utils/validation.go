package utils

import (
	"errors"
	"regexp"
)

// ValidateEmail ensures the email follows a valid format.
func ValidateEmail(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(regex, email); !matched {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidateUsername ensures the username meets length and character constraints.
func ValidateUsername(username string) error {
	if len(username) < 3 || len(username) > 50 {
		return errors.New("username must be between 3 and 50 characters")
	}
	regex := `^[a-zA-Z0-9_]+$`
	if matched, _ := regexp.MatchString(regex, username); !matched {
		return errors.New("username can only contain letters, numbers, and underscores")
	}
	return nil
}

// ValidateAuth0ID ensures the Auth0 ID is not empty.
func ValidateAuth0ID(auth0ID string) error {
	if auth0ID == "" {
		return errors.New("auth0_id cannot be empty")
	}
	return nil
}
