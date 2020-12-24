package main

import (
	"fmt"
	"log"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println(quote.Go())

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message.
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

}