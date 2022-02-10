package greetings

import (
	"fmt"
	"strings"
)

type Conversation interface {
	SayHello() string
	SayGoodbye() string
}

type politeConversation struct {
	name string
}

func NewPoliteConversation(name string) (*politeConversation, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	return &politeConversation{name}, nil
}

func (c politeConversation) SayHello() string {
	return fmt.Sprintf("Good morrow, fair %v!", c.name)
}

func (c politeConversation) SayGoodbye() string {
	return fmt.Sprintf("Cheerio, taa-taa %v, my good lad!", c.name)
}

type rudeConversation struct {
	name string
}

func NewRudeConversation(name string) (*rudeConversation, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	return &rudeConversation{name}, nil
}

func (c rudeConversation) SayHello() string {
	return fmt.Sprintf("What do you want %v?", c.name)
}

func (c rudeConversation) SayGoodbye() string {
	return fmt.Sprintf("Go away %v!", c.name)
}

type ErrBoring struct{}

func (e ErrBoring) Error() string {
	return "come on, that's boring"
}

type ErrEmptyName struct{}

func (e ErrEmptyName) Error() string {
	return "empty name"
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
