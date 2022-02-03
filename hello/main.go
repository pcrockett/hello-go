package main

import (
	"log"

	"philcrockett.com/greetings"
)

func main() {
	log.SetFlags(0) // Disable date / time / etc. from log output

	name := "phil"

	greeting, err := greetings.Hello(name)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(greeting)

	greeting, err = greetings.Goodbye(name)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(greeting)
}
