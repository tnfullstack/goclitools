package main

// Todo1 app is an upgrade todo app with flag options

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo1"
)

func main() {
	const todoFileName = ".todo.json"

	// Parsing command line flags
	task := flag.String("t", "", "Task to be added to the Todo list")
	list := flag.Bool("l", false, "List all tasks")
	complete := flag.Int("c", 0, "Item to mark as completed")
	delete := flag.Int("d", 0, "Item to be deleted")
	flag.Parse()

	// Define an item list
	l := &todo1.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// For no extra arguments, print the list
	case *list:
		// List current todo items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
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
		fmt.Println("Invalid option")
		fmt.Println("usage: [-t | -l | -d]")
		fmt.Println("-t : -t Task name (Add a task)")
		fmt.Println("-l : print the todo list")
		fmt.Println("-d : -d 2 (Delete task #2)")
		fmt.Println("-c : -c 1 (Mark task #1 as Done)")
	}
}
