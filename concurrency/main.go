package main

import (
	"fmt"

	"github.com/adamGrgic/gosandbox/pipes/channelSB"
)

// // Stage 1: Generate numbers and send them to the output channel
// func generate(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// // Stage 2: Square each number received from the input channel
// func square(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// // Stage 3: Print each squared number received from the input channel
// func print(in <-chan int) {
// 	for n := range in {
// 		fmt.Println(n)
// 	}
// }

// func generateNumbersSlice(numberCount int, min int, max int) {
// 	for i := 0; i < numberCount; i++ {
// 		randomNum := rand.Intn(max-min+1) + min
// 	}
// }

func main() {
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
