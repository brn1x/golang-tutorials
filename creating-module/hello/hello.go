package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings func: ")
	log.SetFlags(0)

	names := []string{"brn1x", "kiss"}
	messages, err := greetings.Hellos(names)
	// message, err := greetings.Hello("brn1x")

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
