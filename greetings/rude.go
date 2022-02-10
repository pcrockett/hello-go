package greetings

import "fmt"

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
