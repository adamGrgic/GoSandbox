package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// CSV file read all at once
	// lines is a [][]string variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func main() {
	fmt.Println("running read csv")

	content, err := readCSVFile("test.csv")
	if err != nil {
		fmt.Errorf("Could not read file: ", err)
	}

	fmt.Println(content)

}
