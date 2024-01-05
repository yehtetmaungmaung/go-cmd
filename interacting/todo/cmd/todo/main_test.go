package main_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = "todo.json"
)

// TestMain is the entry point for running tests and cleaning up afterwards.
//
// It takes a single parameter `m` of type `*testing.M`, representing the test suite to be executed.
// It builds the tool, runs the tests, and cleans up afterwards.
//
// Returns the result of the test suite execution.
func TestMain(m *testing.M) {
	// Print "Building tool..." to the console
	fmt.Println("Building tool...")

	// Check the operating system and append ".exe" to `binName` if running on Windows
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// Execute the command "go build -o {binName}" to build the tool
	build := exec.Command("go", "build", "-o", binName)

	// If the build fails, print an error message to `os.Stderr` and exit with a status code of 1
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %q: %s", binName, err)
		os.Exit(1)
	}

	// Print "Running tests..." to the console
	fmt.Println("Running tests...")

	// Execute the tests in the test suite `m` and store the result
	result := m.Run()

	// Print "Cleaning up..." to the console
	fmt.Println("Cleaning up...")

	// Remove the binary file `binName` and the file `fileName`
	os.Remove(binName)
	os.Remove(fileName)

	// Exit with the result of the test suite execution
	os.Exit(result)
}

// TestTodoCLI is a test function for the TodoCLI function.
//
// The function tests the TodoCLI function by performing the following steps:
// - Sets up a test task.
// - Retrieves the current working directory.
// - Constructs the command path.
// - Adds a new task using the TodoCLI function.
// - Lists all the tasks using the TodoCLI function.
// - Compares the expected output with the actual output.
//
// Parameters:
// - t: A testing.T object used for testing.
//
// Return type: None.
func TestTodoCLI(t *testing.T) {
	// Set up a test task.
	task := "test task number 1"
	// Retrieve the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// Construct the command path.
	cmdPath := filepath.Join(dir, binName)
	// Add a new task using the TodoCLI function.
	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	task2 := "test task number 2"
	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdin, task2)
		cmdStdin.Close()
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	// List all the tasks using the TodoCLI function.
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		log.Println(string(out))
		if err != nil {
			t.Fatal(err)
		}

		// Compare the expected output with the actual output.
		expected := fmt.Sprintf("  1: %s\n  2: %s\n", task, task2)
		if expected != string(out) {
			t.Fatalf("expected %q, got %q instead", expected, string(out))
		}
	})
}
