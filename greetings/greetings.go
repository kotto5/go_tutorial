package greetings

import (
	"errors"
	"fmt"
)

// Hello return a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// if a name was received, return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
