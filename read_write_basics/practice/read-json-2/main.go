package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Crop struct {
	Name   string `json:"name"`
	Strain string `json:"strain"`
}

func main() {

	var crops []Crop
	_, err := os.Stat("test.json")
	if err != nil {
		fmt.Println("file does not exist")
	}

	o, err := os.Open("test.json")
	if err != nil {
		fmt.Println("could not open file")
	}
	defer o.Close()

	bytes, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("could not read file")
	}

	json.Unmarshal(bytes, &crops)

	fmt.Println(crops)

}
