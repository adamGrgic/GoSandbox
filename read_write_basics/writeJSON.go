package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func appendToJSONFile(filename string, newPeople []Person) error {
	var people []Person

	// If the file exists, load and unmarshal the existing data
	if fileExists(filename) {
		data, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		if len(data) > 0 {
			if err := json.Unmarshal(data, &people); err != nil {
				return err
			}
		}
	}

	// Append the new people to the slice
	people = append(people, newPeople...)

	// Marshal the updated slice and write back to file
	finalData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, finalData, 0644)
}

func main() {
	newPeople := []Person{
		{"Alice", 30, "alice@example.com"},
		{"Bob", 25, "bob@example.com"},
	}

	if err := appendToJSONFile("people.json", newPeople); err != nil {
		panic(err)
	}
}
