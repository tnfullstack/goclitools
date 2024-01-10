package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Verify and parse arguments
	op := flag.String("op", "sum", "Operation to be executed")
	col := flag.Int("col", 1, "CVS column on which to execute operation")
	flag.Parse()

	if err := run(flag.Args(), *op, *col, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run
func run(filenames []string, op string, col int, out io.Writer) error {
	var opFunc statsFunc

	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if col < 1 {
		return fmt.Errorf("%w, %d", ErrInvalidColumn, col)
	}

	// Validate the operation and define the opFunc accordingly
	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w, %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)

	// Loop through all files adding their data to consolidate
	for _, fn := range filenames {
		// Open the file for reading
		f, err := os.Open(fn)
		if err != nil {
			return fmt.Errorf("cannot open file: %w", err)
		}

		// Parse the CVS into a sclise of float64, numbers
		data, err := csv2float(f, col)
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		// Append the data to consolidate
		consolidate = append(consolidate, data...)
	}
	_, err := fmt.Fprintln(out, opFunc(consolidate))
	return err
}
