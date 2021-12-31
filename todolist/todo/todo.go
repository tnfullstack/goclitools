package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Stringer interface {
	String(v, o bool) string
}

// item struct
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of items
type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

//String prints out a formatted list​
//Implements the fmt.Stringer interface​
func (l List) String(v, o bool) string {
	formatted := ""

	for k, t := range l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		switch {
		case v:
			formatted += fmt.Sprintf("%s%d: %s, date: %v\n", prefix, k+1, t.Task, t.CreatedAt)
		case o:
			if t.Done {
				continue
			}
			formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		default:
			formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		}
	}
	return formatted
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
