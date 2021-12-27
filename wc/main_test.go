package main

import (
	"bytes"
	"testing"
)

// TestCountWords test  the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1\n word2\n word3\n word4\n")
	// fmt.Println(b)
	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("Expcted %d, go %d instead.\n", exp, res)
	}
}

// TestCountLines test the count function set to cound line
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1\n word2\n word3\n word4\n")

	exp := 4

	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
