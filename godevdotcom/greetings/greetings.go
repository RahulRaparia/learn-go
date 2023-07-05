package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func init() {
	currentTime := time.Now().UnixNano()
	fmt.Println("Current time: ", currentTime)
	rand.Seed(currentTime)
}
//Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//If no name -> return error w/message
	if name == ""{
		return "", errors.New("empty name")
	}
	//Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}