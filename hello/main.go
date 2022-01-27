package main

import (
	"fmt"
	"os"

	"philcrockett.com/greetings"
)

func main() {
	name := "foo"

	greeting, err := greetings.Hello(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	greeting, err = greetings.Goodbye(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
