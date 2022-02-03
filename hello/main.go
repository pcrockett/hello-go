package main

import (
	"fmt"
	"log"

	"philcrockett.com/greetings"
)

func main() {
	log.SetFlags(0) // Disable date / time / etc. from log output

	name := "phil"
	conversationType := "pirate"

	var conversation greetings.IConversation
	var err error

	if conversationType == "polite" {
		conversation, err = greetings.NewPoliteConversation(name)
	} else if conversationType == "rude" {
		conversation, err = greetings.NewRudeConversation(name)
	} else if conversationType == "pirate" {
		conversation = PirateConversation{name}
	} else {
		log.Fatalf("Unrecognized conversation type: %v", conversationType)
	}

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(conversation.SayHello())
	log.Println(conversation.SayGoodbye())

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
