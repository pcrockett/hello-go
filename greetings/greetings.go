package greetings

import (
	"errors"
	"fmt"
	"strings"
)

func Hello(name string) (string, error) {
	err := validateName(name)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello %v!", name), nil
}

func Goodbye(name string) (string, error) {
	err := validateName(name)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Goodbye %v!", name), nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("empty name")
	}

	if strings.ToLower(name) == "world" {
		return errors.New("come on, that's boring")
	}

	return nil
}
