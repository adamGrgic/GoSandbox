package channelSB

import (
	"fmt"
)

// Stage 1: Generate numbers and send them to the output channel
func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Stage 2: Square each number received from the input channel
func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Stage 3: Print each squared number received from the input channel
func Print(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func GenerateNumbersChannel(numberCount int, min int, max int) <-chan int {
	fmt.Println("Entering GenerateNumbersChannel")
	randomNum := 5
	out := make(chan int)

	// rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	go func() {
		defer close(out) // Ensure the channel is closed when the goroutine completes
		for i := 0; i < numberCount; i++ {

			fmt.Printf("New random number generated: %d, iteration: %d\n", randomNum, i)
			out <- randomNum // Send the random number to the channel
		}
		fmt.Println("Numbers channel is closed out")
	}()

	return out
}
