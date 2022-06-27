// Declare a main package
package main

// Import the fmt, greetings and rsc.io/quote package
import (
	"fmt"

	"github.com/DanNduati/Go_learn/greetings"
	"rsc.io/quote"
)

// The main function executes by default when you run the main package
func main() {
	//get a greeting message and print it
	message := greetings.Hello("Daniel")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
