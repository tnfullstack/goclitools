package todo_test

import (
	"fmt"
	"goclitools/todolist/todo"
	"testing"
)

// TestAdd()
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	for i := 0; i < 5; i++ {
		l.Add(taskName)
	}

	for i := range l {
		if l[i].Task != taskName {
			t.Errorf("expected %d get #d: %s get %s instead\n", i, taskName, l[i].Task)
		}
		fmt.Println(l[i].Task)
	}
}

// TestComplete
func TestComplete(t *testing.T) {
	l := todo.List{}
	var str []string

	// Add 5 task to the list
	for i := 0; i <= 5; i++ {
		taskName := fmt.Sprintf("New Task %d", i)
		l.Add(taskName)
		str = append(str, taskName)
		fmt.Println("From []str ->", str[i])
	}

	if l[1].Done {
		t.Errorf("tash should not be completed.")
	}

	// Test complete items
	l.Complete(2)

	if !l[1].Done {
		t.Errorf("selected tast %v should be completed\n", l[1].Done)
	}

}
