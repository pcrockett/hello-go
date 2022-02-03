package greetings

import (
	"fmt"
	"strings"
)

type IConversation interface {
	SayHello() string
	SayGoodbye() string
}

type PoliteConversation struct {
	name string
}

func NewPoliteConversation(name string) (*PoliteConversation, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	return &PoliteConversation{name}, nil
}

func (c PoliteConversation) SayHello() string {
	return fmt.Sprintf("Good morrow, fair %v!", c.name)
}

func (c PoliteConversation) SayGoodbye() string {
	return fmt.Sprintf("Cheerio, taa-taa %v, my good lad!", c.name)
}

type RudeConversation struct {
	name string
}

func NewRudeConversation(name string) (*RudeConversation, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	return &RudeConversation{name}, nil
}

func (c RudeConversation) SayHello() string {
	return fmt.Sprintf("What do you want %v?", c.name)
}

func (c RudeConversation) SayGoodbye() string {
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
