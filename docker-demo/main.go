package main

import (
	"fmt"
	"os"
)

func main() {
	// Use container path instead of Windows path
	filePath := "/foo/test.txt"

	// Create or open the file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write a message to the file
	message := "Hello from Go inside a container!"
	_, err = file.WriteString(message)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	file.Sync()
	fmt.Println("Message written to:", filePath)

	return
}
