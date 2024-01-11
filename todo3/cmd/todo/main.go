package main

// Todo3 app is an upgrade from todo2 app with add feature for captureing input from Stdin

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo3"
)

func main() {
	// default file name
	var todoFileName = ".todo.json"

	// Check if the user defined the ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// Parsing command line flags
	add := flag.Bool("a", false, "Task to be added to the Todo list")
	list := flag.Bool("l", false, "List all tasks")
	complete := flag.Int("c", 0, "Item to mark as completed")
	del := flag.Int("d", 0, "Item to be deleted")
	flag.Parse()

	// Define an item list
	l := &todo3.List{}

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

	case *add:
		// When any arguments (excluding flags) are provided, they will be used as a new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Add the task
		l.Add(t)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		if err := l.Delete(*del); err != nil {
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

// getTask function decides where to get the description for a new task from: arguments of STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task connot be blank")
	}

	return s.Text(), nil
}
