package main

import (
	"fmt"
	"log"

	"github.com/rebbyv/go-module-tutorial.git/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	message, err := greetings.Hello("Rebekah")

	if err != nil {
		// If an error was returned, print it to the console and exit the program.
		log.Fatal(err)
	}

	fmt.Println(message)
}