package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// config struct
type config struct {
	ext  string    // filter by file extension
	size int64     // filter by file minimum file size
	list bool      // listing files
	del  bool      // delete files
	wLog io.Writer // write log
}

// program entry
func main() {
	// Parsing commend line flags
	root := flag.String("root", ".", "Root directory to start")
	logFile := flag.String("log", "", "Log delete to this file")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	del := flag.Bool("del", false, "Delete files")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)

	// Intentiate config struct
	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
		wLog: f,
	}

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer f.Close()
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run
func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)
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

		// Delete Files
		if cfg.del {
			return delFile(path, delLogger)
		}

		// List is the default option if nothing else was set
		return listFile(path, out)
	})
}
