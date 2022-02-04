package main

import (
	"clitools/todov1.0/todo"
	"flag"
	"fmt"
	"os"
)

// Hardcoding the file name
const todoFilename = "data/todo.json"

func main() {
	// Parse command line flags
	task := flag.String("t", "", "Adding task to ToDo list")
	list := flag.Bool("l", false, "List ToDo items")
	done := flag.Int("d", 0, "Mark ToDo item as completed")
	flag.Parse()

	// Define a items list
	l := &todo.List{}

	// Get command to read todo items from file
	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments prvovided
	switch {
	// Print the list when no arguments
	case *list:
		// List current todo items
		fmt.Print(l)
		// Add todo item to the list
	case *task != "":
		// Joint the command argument into on string
		// item := strings.Join(os.Args[2:], " ")
		l.Add(*task)
		// Save the new list
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *done > 0:
		// Complete the given item
		if err := l.Complete(*done); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the list
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag prvovided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
