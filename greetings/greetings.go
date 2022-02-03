package greetings

import (
	"fmt"
	"strings"
)

type Conversation struct {
	hello   string
	goodbye string
}

type IConversation interface {
	SayHello() string
	SayGoodbye() string
}

func (c Conversation) SayHello() string {
	return c.hello
}

func (c Conversation) SayGoodbye() string {
	return c.goodbye
}

type ErrBoring struct{}

func (e ErrBoring) Error() string {
	return "come on, that's boring"
}

type ErrEmptyName struct{}

func (e ErrEmptyName) Error() string {
	return "empty name"
}

func Rude(name string) (IConversation, error) {
	err := validateName(name)
	if err != nil {
		return Conversation{}, err
	}

	return Conversation{
		hello:   fmt.Sprintf("What do you want %v?", name),
		goodbye: fmt.Sprintf("Go away %v!", name),
	}, nil
}

func Polite(name string) (IConversation, error) {
	err := validateName(name)
	if err != nil {
		return Conversation{}, err
	}

	return Conversation{
		hello:   fmt.Sprintf("Good morrow, fair %v!", name),
		goodbye: fmt.Sprintf("Cheerio, taa-taa %v, my good lad!", name),
	}, nil
}

func validateName(name string) error {
	if name == "" {
		return ErrEmptyName{}
	}

	if strings.ToLower(name) == "world" {
		return ErrBoring{}
	}

	return nil
}
