package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

const (
	inputFile  = "testdata/test.md"
	resultFile = "test.md.html"
	outputFile = "testdata/test.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result := parseContent(input)

	expected, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("output:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Errorf("result content does not match output file")
	}
}

func TestRun(t *testing.T) {
	if err := run(inputFile); err != nil {
		t.Fatal(err)
	}

	result, err := ioutil.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile(outputFile)
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
