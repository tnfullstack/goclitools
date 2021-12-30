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

// List represents a litt of items
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

// Complete method makes an item as completed by setting done = true and completedAt = current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Add item for 0 base index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes an item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for the zero base index
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save method encode the list in JSON and save it using the provided name
func (l *List) Save(fn string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(fn, js, 0644)
}

// Get method open the provided filename, decodes the JSON data and parse it into a list
func (l *List) Get(fn string) error {
	f, err := os.ReadFile(fn)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(f) == 0 {
		return nil
	}

	return json.Unmarshal(f, l)

}
