package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of items
type List []item

// Add creates a new todo item and append it to the List
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete method marks an item as completed by setting done = true and completedAt = current time
func (l *List) Complete(i int) error {
	list := *l
	if i <= 0 || i > len(list) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Add item for 0 base index
	list[i-1].Done = true
	list[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes an item from the list
func (l *List) Delete(i int) error {
	list := *l
	if i <= 0 || i > len(list) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for the zero base index
	*l = append(list[:i-1], list[i:]...)
	return nil
}

// Save method encode the list in JSON and save it using the provided name
func (l *List) Save(fn string) error {
	js, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(fn, js, 0666)
}

// Get method open the provided filename, decodes the JSON data and parse it into a list
func (l *List) Get(fn string) error {
	file, err := os.ReadFile(fn)
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
