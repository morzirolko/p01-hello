package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the single named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received, return a value that embeds the name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names.
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

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Ujcgjlm is watching, %v. Beware.",
	}
	return formats[rand.Intn(len(formats))]
}
