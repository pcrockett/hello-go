package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}

func Goodbye(name string) string {
	return fmt.Sprintf("Goodbye %v!", name)
}
