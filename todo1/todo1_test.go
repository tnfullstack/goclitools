package todo1

import (
	"fmt"
	"os"
	"testing"
)

// TestAdd tests the ad method of the List type
func TestAdd(t *testing.T) {
	l := List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %s, got %s instead\n", taskName, l[0].Task)
	}
}

// TestComplete test the Complete method of the List type
func TestComplete(t *testing.T) {
	l := List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.\n", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("expect %v, but got %v instead\n", false, l[0].Done)
	}

	// Set the new task to complete (l[0].Done = true)
	l.Complete(1) // l[1-1].Done = true

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

// TestDelete test the Delete method of the List type
func TestDelete(t *testing.T) {
	l := List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
		"New Task 4",
		"New Task 5",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q instead\n", tasks[0], l[0].Task)
	} else {
		fmt.Printf("expected %q, got %q\n", tasks[0], l[0].Task)
	}

	// Delete task at index second index
	l.Delete(2) // l[2-1]
	exp := len(tasks) - 1
	res := len(l)
	if exp != res {
		t.Errorf("expected %d, got %d instead\n", exp, res)
	} else {
		fmt.Printf("expect %d, got %d\n", exp, res)
		for _, v := range l {
			fmt.Printf("%v\n", v)
		}
	}
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := List{}
	l2 := List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %q, go %q instead.\n", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s\n", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s\n", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s\n", err)
	}

	if l1[0].Task != l1[0].Task {
		t.Errorf("task %q should match %q task.\n", l1[0].Task, l2[0].Task)
	}

}
