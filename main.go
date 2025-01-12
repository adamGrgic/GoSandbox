// idea for end goal: create a cli that lets you browse various go modules
// (available in the subdirectories of this project)
// - could also be portable to other systems
// -


// alot of this will be notes from Mastering Go by
// see notes.txt for more verbose breakdowns

package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting up GO Sandbox. Hope you're ready for all the incoming education!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    // get user input
    fmt.Println("Please enter your name")
    var name string

    fmt.Scanln(&name)
    fmt.Println("Your name is",name)
}
