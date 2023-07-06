package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	currentTime := time.Now().UnixNano()
	fmt.Println("Current time: ", currentTime)
	rand.Seed(currentTime)
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//If no name -> return error w/message
	if name == "" {
		return "", errors.New("empty name")
	}
	//Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())
	return message, nil
}

//Hellos retuns a map that associates each of the named people with a greeting message.

func Hellos(names []string) (map[string]string, error) {
	//A map to associate names with messages.
	messages := make(map[string]string)
	//Loop through the received slice of names calling the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"What Upp %v!",
		"Hello, %v. How do you do?",
		"Hey, %v. How are you?",
		"Yo, %v. What's up?",
		"Hi, %v. How's it going?",
		"Hello, %v. How are you?",
		"Hey, %v. What's new?",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
