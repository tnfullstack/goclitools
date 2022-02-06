package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// config struct
type config struct {
	root string
	ext  string
	size int64
	list bool
}

// program entry
func main() {
	// Parsing commend line flags
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	// Intentiate config struct
	c := config{
		root: *root,
		ext:  *ext,
		size: *size,
		list: *list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run
func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}

		// If list was explicitly set, don't do anything else
		if cfg.list {
			return listFile(path, out)
		}

		// List is the default option if nothing else was set
		return listFile(path, out)
	})
}
