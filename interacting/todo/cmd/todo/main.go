package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/yehtetmaungmaung/interacting/todo"
)

var todoFileName = "todo.json"

// main is the entry point of the program.
//
// It reads a todo list from a file, adds a task to the list, and saves the updated list back to the file.
// The todo list is printed to the console if there are no command line arguments.
//
// Parameters:
// None.
//
// Return type:
// None.
func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for the Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2024\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information: %s [add|list|complete] [task]\n", os.Args[0])
		flag.PrintDefaults()
	}
	add := flag.Bool("add", false, "add task to the ToDo list")
	list := flag.Bool("list", false, "list all tasks")
	complete := flag.Int("complete", 0, "item to be completed")
	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	// Create an instance of the todo list
	l := &todo.List{}

	// Read the todo list from a file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Check if there are any command line arguments
	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}

// getTask reads a task from the provided reader or returns the first argument.
// If no arguments are provided, it reads a line from the reader and returns it.
// If the line is blank, it returns an error.
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		// If arguments are provided, return the first argument
		return args[0], nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		// If there was an error scanning the input, return the error
		return "", err
	}

	if len(s.Text()) == 0 {
		// If the scanned line is blank, return an error
		return "", fmt.Errorf("task cannot be blank")
	}

	// Return the scanned line
	return s.Text(), nil
}
