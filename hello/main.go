package main

import (
	"log"

	"philcrockett.com/greetings"
)

func main() {
	log.SetFlags(0) // Disable date / time / etc. from log output

	name := "phil"
	conversationType := "polite"

	var conversation greetings.IConversation
	var err error

	if conversationType == "polite" {
		conversation, err = greetings.NewPoliteConversation(name)
	} else if conversationType == "rude" {
		conversation, err = greetings.NewRudeConversation(name)
	} else {
		log.Fatalf("Unrecognized conversation type: %v", conversationType)
	}

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(conversation.SayHello())
	log.Println(conversation.SayGoodbye())

}
