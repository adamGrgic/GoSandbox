package main

import (
	"fmt"
	"os"
	"strings"

	todo "github.com/adamGrgic/GoSandbox/cli-app-1"
)

const todoFileName = ".todo.json"

func main() {

	// Define an items list
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	}

}
