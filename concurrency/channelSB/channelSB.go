package channelSB

import (
	"fmt"
	"math/rand"
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
	out := make(chan int)
	go func() {
		for i := 0; i < numberCount; i++ {
			randomNum := rand.Intn(max-min+1) + min
			fmt.Printf("New random number generated: %d, iteration: %d", randomNum, i)

		}
		fmt.Printf("Numbers channel is closed out")
		close(out)
	}()
	return out
}
