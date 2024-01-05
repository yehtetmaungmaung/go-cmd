package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task       string
	Done       bool
	CreatedAt  time.Time
	CompleteAt time.Time
}

type List []item

// Add adds a new task to the list.
//
// The task parameter is a string representing the task to be added.
// This function does not return any value.
func (l *List) Add(task string) {
	t := item{
		Task:       task,
		Done:       false,
		CreatedAt:  time.Now(),
		CompleteAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete marks the item at the given index as complete.
//
// It takes an integer parameter `i` representing the index of the item to mark as complete.
// It returns an error.
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompleteAt = time.Now()

	return nil
}

// Delete deletes an item from the List at the specified index.
//
// Parameters:
// - i: the index of the item to be deleted
//
// Returns:
// - error: if the index is out of range
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save saves the List to a file.
//
// It takes a filename string as a parameter and returns an error.
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get retrieves a file from the List.
//
// It takes a filename as a parameter and returns an error.
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
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

// String returns a formatted string representation of the List.
//
// It iterates over each element in the List and constructs a formatted string representation
// by concatenating the index, task status, and task description.
//
// Parameters:
// - None
//
// Return type:
// - string
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "x "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
