package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
    fmt.Println("Hello world!")

    flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for the Pragmatic Bookshelf\n", os.Args[0])

		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "list all tasks")
    todos := flag.Bool("todos", false, "list all uncompleted tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

    switch {
	case *list:
        fmt.Println("the list flag was given!")
	case *add:
        fmt.Println("the add flag was given!")
	case *complete > 0:
        fmt.Println("the complete flag was given!")
    case *todos:
        fmt.Println("the todos flag was given!")
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}


}
