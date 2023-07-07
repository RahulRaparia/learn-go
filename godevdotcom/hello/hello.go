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

	// A slice of names."
	names := []string{"Cici", "Fifi", "Hector", "Pipi"}

	// Request a greeting for each name in the names slice.
	messages, err := greetings.Hellos(names)
	// If an error was returned, printing it to the console and esit the program
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of messages to the console
	fmt.Println(messages)

	//write code to print output on console in new line 
	// fmt.Println("Hello, World!")
}
