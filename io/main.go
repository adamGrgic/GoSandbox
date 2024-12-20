package main

import (
	"bufio"
	"fmt"
	"os"
    "log"
)


func main() {
    fmt.Println("running io module...")

    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {

        fmt.Println(scanner.Err())

        fmt.Println(scanner.Bytes())




        line := scanner.Text()
        fmt.Println(line)

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
