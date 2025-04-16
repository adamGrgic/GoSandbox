package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
	records := [][]string{
		{"Name", "Age", "Email"},
		{"Alice", "30", "alice@example.com"},
		{"Bob", "25", "bob@example.com"},
		{"Charlie", "35", "charlie@example.com"},
	}

	err := writeCSVFile("testwriter.csv", records)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}

	fmt.Println("CSV file written successfully!")
}
