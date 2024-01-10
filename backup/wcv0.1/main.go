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
	// Calling the count function to count the number of words
	// received from the Standard Input and print it out
	input := os.Stdin

	fmt.Println("Word count =", count(input))
}

// count
func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wc := 0

	// For every word scanned, increement the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc

}
