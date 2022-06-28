// Declare a main package
package main

// Import the fmt, greetings and rsc.io/quote package
import (
	"fmt"
	"log"

	"github.com/DanNduati/Go_learn/greetings"
	"rsc.io/quote"
)

// The main function executes by default when you run the main package
func main() {
	// set properties of the predefined logger,including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	// create a slice of names
	names := []string{"Daniel", "Nduati", "Chege"}

	//get a greeting message and print it
	messages, err := greetings.Hellos(names)
	if err != nil {
		//Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(err)
	}
	// if no error was returned print the returned message to the console
	for name, message := range messages {
		fmt.Println("Greeting for ", name, ":", message)
	}

	fmt.Println(quote.Go())
}
