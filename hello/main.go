package main

import (
	"log"

	"philcrockett.com/greetings"
)

func main() {
	log.SetFlags(0) // Disable date / time / etc. from log output

	name := "phil"

	conversation, err := greetings.Polite(name)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(conversation.Hello)
	log.Println(conversation.Goodbye)

}
