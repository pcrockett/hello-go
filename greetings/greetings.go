package greetings

import (
	"errors"
	"fmt"
	"strings"
)

type Conversation struct {
	Hello   string
	Goodbye string
}

func Rude(name string) (Conversation, error) {
	err := validateName(name)
	if err != nil {
		return Conversation{}, err
	}

	return Conversation{
		Hello:   fmt.Sprintf("What do you want %v?", name),
		Goodbye: fmt.Sprintf("Go away %v!", name),
	}, nil
}

func Polite(name string) (Conversation, error) {
	err := validateName(name)
	if err != nil {
		return Conversation{}, err
	}

	return Conversation{
		Hello:   fmt.Sprintf("Greetings, dear %v!", name),
		Goodbye: fmt.Sprintf("I bid you farewell, %v!", name),
	}, nil
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
