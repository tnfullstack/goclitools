package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestWordsCount(t *testing.T) {
	b := bytes.NewBufferString("Word1 word2 word3 word4\n")
	exp := 4
	res := wordsCount(b)
	if res != exp {
		t.Errorf("expected %d, got %d instead.\n", exp, res)
	}
}

func TestLinesCount(t *testing.T) {
	b := bytes.NewBufferString("word1\nword2\nword3\nword4\n")
	exp := 4
	res := linesCount(b)
	if res != exp {
		t.Errorf("expected %d, got %d instead.\n", exp, res)
	}
}

func TestBytesCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 24
	res := bytesCount(b)
	if res != exp {
		t.Errorf("expected %d, got %d instead.\n", exp, res)
	}
}
