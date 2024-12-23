package main

import(
    "fmt"
    "bufio"
    "io"
    "os"
    "flag"
)

func count(r io.Reader, countLines bool) int {

    // A scanner is used to read text from a Reader (such as files)
    scanner := bufio.NewScanner(r)

    // If the count lines flag is not set, we want to count words so we define
    // Define the scanner split type to words (default is split by lines)
    if !countLines {
        scanner.Split(bufio.ScanWords)
    }
    // Defining a counter
    wc := 0

    for scanner.Scan() {
        wc++
    }

    return wc
}

func main() {
    fmt.Println("starting up cli application")

    // Defining a Boolean flag -l to count lines instead of words
    lines := flag.Bool("l", false, "Count lines")

    // Parsing the flags provided by the user
    flag.Parse()



    fmt.Println(count(os.Stdin, *lines))
}
