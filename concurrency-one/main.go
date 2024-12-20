// module for testing concurrent operations
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// // Section 1
	// // Simple Number Example
	// // Step 1: Start the pipeline by generating numbers
	// fmt.Println("Generating numbers...")

	// numbers := channelSB.GenerateNumbersChannel(100, 123, 250)

	// fmt.Println("Generated numbers:")
	// for value := range numbers {
	// 	fmt.Printf("Number: %d\n", value)
	// }

	// // Step 2: Pass the generated numbers to the square stage
	// squares := channelSB.Square(numbers)

	// // Step 3: Print the squared numbers
	// print(squares)

	// // Section 2
	// // Create a buffer pipe (a pipe where the buffer size is sure to be reached)

	var wg sync.WaitGroup
	maxConcurrent := 10 // Number of concurrent goroutines allowed
	semaphore := make(chan struct{}, maxConcurrent)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Acquire the semaphore slot
			semaphore <- struct{}{}

			// Critical section
			fmt.Printf("Goroutine %d is running\n", id)
			time.Sleep(2 * time.Second) // Simulate some work

			// Release the semaphore slot
			<-semaphore

			fmt.Printf("Goroutine %d is done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines have completed.")
}
