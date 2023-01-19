package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings func: ")
	log.SetFlags(0)

	message, err := greetings.Hello("brn1x")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
