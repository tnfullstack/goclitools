package main

import (
	"bytes"
	"testing"
)

// TestCountLines test the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("Word1 word2 word3 word4\nline2\nline3 word1")

	exp := 3 // 3 lines

	res := count(b, true)

	if res != exp {
		t.Errorf("expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {

	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("expected %d, got %d instead.\n", exp, res)
	}
}
