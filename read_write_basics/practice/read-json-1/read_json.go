package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type People struct {
	Name              string `json: "name"`
	Role              string `json: "role"`
	SecurityClearance int    `json: "security_clearance"`
}

func main() {
	var people []People

	_, err := os.Stat("test.json")
	if err != nil {
		fmt.Println("file does not exist")
		os.Exit(1)
	}

	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Error opening file: ", file)
	}
	defer file.Close()

	bytes, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Errorf("could not read file ", err)
	}

	json.Unmarshal(bytes, &people)

	fmt.Println(people)
}
