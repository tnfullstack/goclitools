package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const prompt = `
Word count (wc) reads the text file an count the
number of words in the file and return the total number.

How to run the program:
$ wc < filename.txt

`

func main() {

	fmt.Println(prompt)
	fmt.Println(count(os.Stdin))

	// TestCountWords()
}

// count
func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)

	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is splist by lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the toal
	return wc
}
