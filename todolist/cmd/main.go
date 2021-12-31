package main

import (
	"flag"
	"fmt"
	"goclitools/todo"
	"os"
)

const usage = `
Usage: 
./todo -list (list todo items)
./todo -task (add new todo item)
./todo -complete 1 (marks item 1 as completed)

`

var todoFileName = "data.json"

func main() {

	// Check if the user definded ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// Parsing command-line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all task")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	// Define an item list
	l := &todo.List{}

	// Use the Get method to read todo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number arguments provided
	switch {
	case *list:
		// list current todo items
		fmt.Println("ToDo List")
		fmt.Println("----------")
		fmt.Print(l)

	case *complete > 0:
		// complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add the task
		l.Add(*task)
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Print(usage)
	}
}
