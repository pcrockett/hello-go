package greetings

import "strings"

func validateName(name string) error {
	if name == "" {
		return ErrEmptyName{}
	}

	if strings.ToLower(name) == "world" {
		return ErrBoring{}
	}

	return nil
}
