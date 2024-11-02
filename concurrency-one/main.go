// module for testing concurrent operations
package main

import (
	"fmt"

	"github.com/adamGrgic/concurrency-one/channelSB"
)

func main() {

	// Simple Number Example
	// Step 1: Start the pipeline by generating numbers
	fmt.Println("Generating numbers...")

	numbers := channelSB.GenerateNumbersChannel(100, 123, 250)

	fmt.Println("Generated numbers:")
	for value := range numbers {
		fmt.Printf("Number: %d\n", value)
	}

	// Step 2: Pass the generated numbers to the square stage
	squares := channelSB.Square(numbers)

	// Step 3: Print the squared numbers
	print(squares)
}
