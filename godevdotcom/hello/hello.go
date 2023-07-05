package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined logger, including the log entry prefix and a flag to disable printing the time, source file, and the line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Request a greeting message.
	message, err := greetings.Hello("ravikanth")
	// If an error was returned, printing it to the console and esit the program
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message to the console
	fmt.Println(message)
}
