package main

import (
	"bytes"
	"fmt"
	"testing"
)

// TestCountWords test  the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	fmt.Println(b)
	exp := 4

	res := count(b)

	if res != exp {
		t.Errorf("Expcted %d, go %d instead.\n", exp, res)
	}
}
