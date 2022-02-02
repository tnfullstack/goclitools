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
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Parsing the flags provided by the user
	flag.Parse()

	input := os.Stdin

	if !*lines {
		// Call linesCount function, then printout the return value
		fmt.Println("Word count =", wordsCount(input))
	} else {
		// Call the linesCount function and printout the return value
		fmt.Println("Words count =", linesCount(input))
	}
}

// linescount
func linesCount(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	lc := 0
	for scanner.Scan() {
		lc++
	}
	return lc
}

// wordsCount
func wordsCount(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	// Defining the counter
	wc := 0
	// For ever word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	return wc
}
