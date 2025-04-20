package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CSVRecord struct {
	Name  string
	Age   int
	Email string
}

func writeCSVFile(filepath string, data [][]string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	data := []CSVRecord{
		{"Alice", 30, "alice@example.com"},
		{"Bob", 25, "bob@example.com"},
		{"Charlie", 35, "charlie@example.com"},
	}

	filename := "testwriter.csv"
	exists := false

	_, err := os.Stat(filename)
	if err == nil {
		exists = true
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("could not open file ", file)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// If file didn't exist before, write headers
	if !exists {
		if err := writer.Write([]string{"Name", "Age", "Email"}); err != nil {
			log.Panic("could not write headers to file")
		}
	}

	// Write records
	for _, p := range data {
		record := []string{p.Name, strconv.Itoa(p.Age), p.Email}
		if err := writer.Write(record); err != nil {
			log.Panic("could not write records to file")
		}
	}

	fmt.Println("CSV file written successfully!")
}
