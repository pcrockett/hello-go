package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"philcrockett.com/greetings"
)

var cliParams struct {
	Name             string `help:"Name of person to whom you want to say hello / goodbye" arg:""`
	ConversationType string `help:"Type of conversation you want to have (polite, rude, pirate)" enum:"polite,rude,pirate" default:"polite" short:"t"`
}

func main() {
	log.SetFlags(0) // Disable date / time / etc. from log output

	kong.Parse(&cliParams)

	conversation, err := createConversation(
		cliParams.ConversationType, cliParams.Name,
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(conversation.SayHello())
	log.Println(conversation.SayGoodbye())

}

type ErrUnrecognizedConversationType struct {
	conversationType string
}

func (e ErrUnrecognizedConversationType) Error() string {
	return fmt.Sprintf("unrecognized conversation type: \"%v\"", e.conversationType)
}

func createConversation(conversationType string, name string) (greetings.IConversation, error) {
	var conversation greetings.IConversation
	var err error

	if conversationType == "polite" {
		conversation, err = greetings.NewPoliteConversation(name)
	} else if conversationType == "rude" {
		conversation, err = greetings.NewRudeConversation(name)
	} else if conversationType == "pirate" {
		conversation = PirateConversation{name}
	} else {
		conversation = nil
		err = ErrUnrecognizedConversationType{conversationType}
	}

	return conversation, err
}

type PirateConversation struct {
	name string
}

func (c PirateConversation) SayHello() string {
	return fmt.Sprintf("Ahoy %v, ye scallywag!", c.name)
}

func (c PirateConversation) SayGoodbye() string {
	return fmt.Sprintf("Yarrg me mateys! %v be walkin the plank!", c.name)
}
