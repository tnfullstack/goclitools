package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// item struct represents a todo item
type Item struct {
	Task       string
	Done       bool
	CreateAt   time.Time
	CompleteAt time.Time
}

// List represents a list of ToDo items
type List []Item

// String prints out a formatted List
// Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for i, item := range *l {
		prefix := "  "
		if item.Done {
			prefix = "X "
		}
		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s %d: %s\n", prefix, i+1, item.Task)
	}
	return formatted
}

// Add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := Item{
		Task:       task,
		Done:       false,
		CreateAt:   time.Now(),
		CompleteAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete method marks a todo item as completed by setting done = true
// and CompleteAt to current Time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	// List = [0][1][2][3] (This is easy visualization only)
	ls[i-1].Done = true
	ls[i-1].CompleteAt = time.Now()

	return nil
}

// Delete method deletes a todo from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0 base index
	// List = [0][1][2][3]
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method saves the list from []List to JSON file
func (l *List) Save(fname string) error {
	js, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(fname, js, 0644)
}

// Get method opens the prvovided file name, decodes the JSON data and parses it into a list
func (l *List) Get(fname string) error {
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
