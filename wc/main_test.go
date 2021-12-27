package main

import (
	"bytes"
	"fmt"
	"testing"
)

// TestCountWords test  the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1\n word2\n word3\n word4\n")
	fmt.Println(b)
	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("Expcted %d, go %d instead.\n", exp, res)
	}
}
