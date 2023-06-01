package validator

import (
	"errors"
	"regexp"
)

func (v *Validator) Password(password string) error {
	var length int
	for _, s := range password {
		length++
		switch s {
		case ' ':
			return errors.New("spaces in the password are not allowed")
		case '%', '$', '#', '*', '/', '\\':
			return errors.New("special characters are not allowed in the password")
		}
	}
	if length < 5 {
		return errors.New("the password must be longer than 4 characters")
	}

	return nil
}

func (v *Validator) Email(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return errors.New("the email is not valid")
	}
	return nil
}
