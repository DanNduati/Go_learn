// Declare a main package
package main

// Import the fmt and rsc.io/quote package
import (
	"fmt"

	"rsc.io/quote"
)

// The main function executes by default when you run the main package
func main() {
	fmt.Println(quote.Go())
}
