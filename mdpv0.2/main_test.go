package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	inputFile  = "testdata/test.md"
	outputFile = "testdata/test.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("output:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("result content does not match output file")
	}
}

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer

	if err := run(inputFile, "", &mockStdOut, true); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mockStdOut.String())
	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("output:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Errorf("result content does not match output file")
	}
	os.Remove(resultFile)
}
