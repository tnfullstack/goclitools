package main

import (
	"fmt"
	"goclitools/todov0.1/todo"
)

func main() {
	l := todo.List{}
	taskName := "New Task"

	// Add 5 task to the list
	for i := 0; i <= 5; i++ {
		str := fmt.Sprintf("%s %d", taskName, i)
		l.Add(str)
		fmt.Println(l[i].Task)
	}
}
