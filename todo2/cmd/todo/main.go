package main

// Todo2 app is an upgrade from todo1 app with feature improving the list output

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo2"
)

func main() {
	// default file name
	var todoFileName = ".todo.json"

	// Check if the user defined the ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// Parsing command line flags
	task := flag.String("t", "", "Task to be added to the Todo list")
	list := flag.Bool("l", false, "List all tasks")
	complete := flag.Int("c", 0, "Item to mark as completed")
	delete := flag.Int("d", 0, "Item to be deleted")
	flag.Parse()

	// Define an item list
	l := &todo2.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	case len(os.Args) == 1:
		// For no extra arguments, print the list
		fmt.Print(l)
	case *list:
		// List current todo items
		fmt.Print(l)
	case *complete > 0:
		// mark the given item as completed
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		// Concatenate all arguments with a space
		item := strings.Join(os.Args[2:], " ")

		// Add the task
		l.Add(item)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		fmt.Println("usage: [-t | -l | -d]")
		fmt.Println("-t : -t Task name (Add a task)")
		fmt.Println("-l : print the todo list")
		fmt.Println("-d : -d 2 (Delete task #2)")
		fmt.Println("-c : -c 1 (Mark task #1 as Done)")
		os.Exit(1)
	}
}
