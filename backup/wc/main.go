package main

import (
	"bufio"
	"flag"
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

	// Calling the counts function to count the number of words (or lines)
	fmt.Print(prompt)

	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")

	// Parsing the flags provided by the user
	flag.Parse()

	var fg string
	if !*lines {
		fg = "words"
	} else {
		fg = "lines"
	}

	ct := count(os.Stdin, *lines)

	fmt.Printf("Number of %s count is %d\n", fg, ct)

}

// count
func count(r io.Reader, countLines bool) int {

	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// if the count lines flag is not set, we want to count the words
	// the scanner split type to words, (dafault is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	} else {
		scanner.Split(bufio.ScanLines)
	}

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the toal
	return wc
}
