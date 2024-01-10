package main

import (
	"clitools/todov0.1/todo"
	"fmt"
	"os"
	"strings"
)

// Hardcoding the file name
const todoFilename = "data/todo.json"

func main() {
	// Define a items list
	l := &todo.List{}

	// Get reads the item from the file
	err := l.Get(todoFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments prvovided
	switch {
	// Print the list when no arguments
	case len(os.Args) == 1:
		// List current todo items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// Concatenate all prvovided arguments with a space and add to
	// the list as an item
	default:
		item := strings.Join(os.Args[1:], " ")
		// add the task
		l.Add(item)
	}

	// Save the new list
	if err := l.Save(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
