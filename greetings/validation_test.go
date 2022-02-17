package greetings

import (
	"errors"
	"testing"
)

func TestEmpty(t *testing.T) {
	var name string
	err := validateName(name)
	if !errors.Is(err, ErrEmptyName{}) {
		t.Fatalf("Empty string didn't return an ErrEmptyName")
	}
}

func TestWorld(t *testing.T) {
	names := []string{
		"world",
		"WoRlD",
		"WORLD",
	}

	for _, n := range names {
		err := validateName(n)
		if !errors.Is(err, ErrBoring{}) {
			t.Fatalf("\"%v\" didn't return an ErrBoring", n)
		}
	}
}

func TestValidName(t *testing.T) {
	names := []string{
		"Kevin",
		"Jeeves",
		"Dread Pirate Roberts",
	}

	for _, n := range names {
		err := validateName(n)
		if err != nil {
			t.Fatalf("\"%v\" returned an error", n)
		}
	}
}
