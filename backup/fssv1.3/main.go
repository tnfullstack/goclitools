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
	arc  string    // archive file
}

// program entry
func main() {
	// Parsing commend line flags
	dir := flag.String("dir", ".", "Root directory to start")
	log := flag.String("log", "", "Log delete to this file")
	list := flag.Bool("list", false, "List files only")
	arc := flag.String("arc", "", "Archive directory")
	del := flag.Bool("del", false, "Delete files")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
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
		arc:  *arc,
	}

	if *log != "" {
		f, err = os.OpenFile(*log, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if err := run(*dir, os.Stdout, c); err != nil {
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

		// Archive files and continue if successful
		if cfg.arc != "" {
			if err := acrchiveFile(root, cfg.arc, path); err != nil {
				return err
			}
		}

		// Delete Files
		if cfg.del {
			return delFile(path, delLogger)
		}

		// List is the default option if nothing else was set
		return listFile(path, out)
	})
}
