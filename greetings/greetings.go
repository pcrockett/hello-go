package greetings

import (
	"errors"
	"fmt"
	"strings"
)

type Conversation struct {
	hello   string
	goodbye string
}

func (c Conversation) SayHello() string {
	return c.hello
}

func (c Conversation) SayGoodbye() string {
	return c.goodbye
}

func Rude(name string) (Conversation, error) {
	err := validateName(name)
	if err != nil {
		return Conversation{}, err
	}

	return Conversation{
		hello:   fmt.Sprintf("What do you want %v?", name),
		goodbye: fmt.Sprintf("Go away %v!", name),
	}, nil
}

func Polite(name string) (Conversation, error) {
	err := validateName(name)
	if err != nil {
		return Conversation{}, err
	}

	return Conversation{
		hello:   fmt.Sprintf("Good morrow, dear %v!", name),
		goodbye: fmt.Sprintf("I bid thee farewell, dear %v!", name),
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
