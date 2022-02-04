package todo_test

import (
	"clitools/todolist/todo"
	"fmt"
	"io/ioutil"
	"os"
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

// TestDelete
func TestDelete(t *testing.T) {
	l := todo.List{}
	var tasks []string
	for i := 0; i <= 5; i++ {
		temp := fmt.Sprintf("New Task %d", i)
		tasks = append(tasks, temp)
		l.Add(temp)
	}

	if l[2].Task != tasks[2] {
		t.Errorf("expected %q, got %q\n", tasks[2], l[2].Task)
	}

	if len(l) == len(tasks) {
		fmt.Printf("expected %d, got %d instead\n", len(l), len(tasks))
	}

	l.Delete(2)

	if len(l) == len(tasks) {
		t.Errorf("expected length %d, got %d instead\n", len(l), len(tasks))
	}
}

// TestSaveGet
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")

	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
