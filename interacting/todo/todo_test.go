package todo_test

import (
	"os"
	"testing"

	"github.com/yehtetmaungmaung/interacting/todo"
)

// TestAdd tests the functionality of the Add method in the List struct.
//
// It creates a new List instance, adds a task to it, and then checks if the task was added successfully.
// The function takes a testing.T object as a parameter and does not return anything.
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l[0].Task)
	}
}

// TestComplete is a Go function that tests the Complete method of the List struct.
//
// The function creates a new List and adds a task to it using the Add method.
// It then asserts that the task was added correctly by comparing its name to the expected value.
// After that, it asserts that the new task is not marked as completed.
// Next, it calls the Complete method on the List to mark the task as completed.
// Finally, it asserts that the task is indeed marked as completed.
func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("new task should not be completed.")
	}

	l.Complete(1)
	if !l[0].Done {
		t.Errorf("task should be completed.")
	}
}

// TestDelete is a unit test for the Delete method of the List struct.
//
// It tests the functionality of deleting an element from the list.
// The test creates a new List instance and adds some tasks to it.
// Then it deletes an element from the list and checks if the length
// of the list and the remaining task in the list are as expected.
func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q instead.", tasks[0], l[0].Task)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("expected list length %d, got %d instead.", 2, len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}

// TestSaveGet tests the Save and Get functions.
//
// It creates two empty todo.List instances, adds a new task to one of them,
// saves the list to a temporary file, retrieves the list from the file, and
// compares the tasks in the two lists.
//
// The function takes *testing.T as a parameter and does not return anything.
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "todo_test")
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
		t.Errorf("task saved and loaded should be the same.")
	}
}
