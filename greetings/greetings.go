package greetings

import "fmt"

// Hello returns a greeting that embeds the name in the message
func Hello(name string) string {
	// Use the fmt package's Sprintf function to create a greeting message.
	// Sprintf substitutes the name parameter's value in the %v format verb.
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}
