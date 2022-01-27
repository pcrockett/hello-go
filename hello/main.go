package main

import (
	"fmt"

	"philcrockett.com/greetings"
)

func main() {
	name := "world"

	greeting, err := greetings.Hello(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(greeting)

	greeting, err = greetings.Goodbye(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(greeting)
}
