package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	inputFile = "./testdata/test1.md"
	// resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result := parseContent(input)

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
}

func TestRun(t *testing.T) {
	var mocStdout bytes.Buffer

	if err := run(inputFile, &mocStdout); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mocStdout.String())

	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("goden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match goden file")
	}

	os.Remove(resultFile)
}
