package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// sum
func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}
	return sum
}

// avg
func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// statsFunc defines a generic statistical function
type statsFunc func(data []float64) float64

// csv2float
func csv2float(r io.Reader, col int) ([]float64, error) {
	// Create the CVS reader used to read data from CVS files
	cr := csv.NewReader(r)

	// Adjusting fo 0 based index
	col--

	// Read in all CVS Data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	var data []float64

	// Loop through all records
	for i, r := range allData {
		if i == 0 {
			continue
		}
		// checking number of columns in CVS file
		if len(r) <= col {
			// File does not have that many columns
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(r))
		}
		// Try to convert data read into a float number
		v, err := strconv.ParseFloat(r[col], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}
	// Return the slice of float64 and nil error
	return data, nil
}
