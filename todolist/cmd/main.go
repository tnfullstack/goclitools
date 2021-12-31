package main

import (
	"flag"
	"fmt"
	"goclitools/todolist/todo"
	"os"
)

const (
	usage = `
Usage: 
./todo -list (list todo items)
./todo -complete 1 (marks item 1 as completed)
./todo -verbose (list todo items with date and time)
./todo -open (list only the open todo)
./todo -add (add new todo item)

`
	head = `
ToDo List
-----------
`
)

var todoFileName = "data.json"

func main() {
	var v, o bool
	// Check if the user definded ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// Parsing command-line flags
	add := flag.Bool("add", false, "Task to the ToDo list")
	// task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	del := flag.Int("delete", 0, "Item to be deleted")
	ver := flag.Bool("verbose", false, "List all tasks with data and time")
	open := flag.Bool("open", false, "List only open task")
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
		v, o = false, false
		s := l.String(v, o)
		// list current todo items
		fmt.Print(head)
		fmt.Print(s)
	case *ver:
		v, o = true, false
		s := l.String(v, o)
		fmt.Print(head)
		fmt.Print(s)
	case *add:
		t, err := todo.GetTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *open:
		v, o = false, true
		s := l.String(v, o)
		fmt.Print(head)
		fmt.Print(s)
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
	case *del > 0:
		// delete a given item
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Print(usage)
	}
}
