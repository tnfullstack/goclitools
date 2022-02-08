package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// executer interface
type executer interface {
	execute() (string, error)
}

// programe entry
func main() {
	// declare command flags
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	// call the run function
	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run
func run(proj string, out io.Writer) error {
	fmt.Println(proj)
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	pipeline := make([]executer, 3)
	// fmt.Println("Pipeline", pipeline)

	pipeline[0] = newStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		proj,
		[]string{"build", ".", "errors"},
	)

	pipeline[1] = newStep(
		"go test",
		"go",
		"Go Test: SUCCESS",
		proj,
		[]string{"test", "-v"},
	)

	pipeline[2] = newExcepStep(
		"go fmt",
		"gofmt",
		"Gofmt: SUCCESS",
		proj,
		[]string{"-l", "."},
	)

	for _, s := range pipeline {
		// fmt.Println("from run loop", s)
		msg, err := s.execute()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
