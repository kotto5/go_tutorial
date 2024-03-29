package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello return a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// if a name was received, return a greeting that embeds the name in a message.
	message := fmt.Sprintf(ramdomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
func ramdomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying
	// a random for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
