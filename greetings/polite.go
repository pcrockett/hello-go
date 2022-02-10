package greetings

import "fmt"

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
