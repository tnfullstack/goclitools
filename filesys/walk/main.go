package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	// extenstion to filter out
	ext string
	// min file size
	size int64
	// list files
	list bool
}

func main() {
	// Parsing command line flags
	root := flag.String("r", ".", "Root directory to start")
	// Action options
	list := flag.Bool("l", false, "List files only")
	// Filter options
	ext := flag.String("e", "", "File extension to filter out")
	size := flag.Int64("s", 0, "Minimum file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}

		// List is the default option if nothing else was set
		return listFile(path, out)
	})
}
