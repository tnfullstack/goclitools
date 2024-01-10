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
	wcount := flag.Bool("w", false, "Count words")
	bcount := flag.Bool("b", false, "Count bytes")
	// Parsing the flags provided by the user
	flag.Parse()

	input := os.Stdin

	if *wcount {
		// Call wordsCount function, then printout the return value
		fmt.Println("words count =", wordsCount(input))
	} else if *bcount {
		// Call the bytesCount function, then printout the return value
		fmt.Println("bytes count =", bytesCount(input))
	} else {
		// Call the linesCount function and printout the return value
		fmt.Println("lines count =", linesCount(input))
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
	// For ever word scanned, increase the counter
	for scanner.Scan() {
		wc++
	}
	return wc
}

// bytesCount
func bytesCount(r io.Reader) int {
	// A canner is used to read test from a Reader (such as files)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)
	bc := 0
	// For ever byte scanned, increase the counter
	for scanner.Scan() {
		bc++
	}
	return bc
}
