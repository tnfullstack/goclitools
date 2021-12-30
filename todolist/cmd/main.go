package main

import (
	"flag"
	"fmt"
	"goclitools/todo"
	"os"
)

const (
	todoFileName = ".todo.json"
	usage        = `
Usage: 
./todo -list (list todo items)
./todo -task (add new todo item)
./todo -complete 1 (marks item 1 as completed)

`
)

func main() {

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
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
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
		fmt.Println("TASK LIST")
		fmt.Println("----------")
		for i, item := range *l {
			if !item.Done {
				fmt.Printf("#%d %s\n", i+1, item.Task)
			}
		}
		fmt.Print(usage)
	}
}
