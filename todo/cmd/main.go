package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tnfullstack/goclitools/todo"
)

const todoFileName = ".todo.json"

func main() {

	// Define an item list
	l := &todo.List{}

	// Use the Get method to read todo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number arguments provided
	switch {
	// For no extra argument print the list
	case len(os.Args) == 1:
		// list current todo items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
