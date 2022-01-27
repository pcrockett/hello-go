package main

import (
	"fmt"

	"philcrockett.com/greetings"
)

func main() {
	fmt.Println(greetings.Hello("foo"))
	fmt.Println(greetings.Goodbye("foo"))
}
