package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
`
	footer = `
	</body>
</html>
`
)

func main() {
	// Parse flags
	fileName := flag.String("file", "", "Markdown file to repview")
	flag.Parse()

	// If user did not provide input file, show usage
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}
	// Call and pass the cmd flag and fileName to run func
	if err := run(*fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run function reads .md file from cli input, pass the input buffer
// to parseContent function for html parsing, receive the html []byte
// then pass []byes and the .html file name to saveHTML for writing to file
func run(fname string) error {
	fmt.Println(fname) // filename just as it type in from comand-line
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(fname)
	fmt.Println("Content of the .md file\n", string(input)) // for testing purpose
	if err != nil {
		return err
	}
	// call and pass the .md []byte to parseContent func for html parsing
	htmlData := parseContent(input)
	fmt.Println(string(htmlData)) // This is just for testing purpose
	// filepath.Base take the filename, cleanup path and only keep the base filename
	// the add the .html to the end of the file name
	outName := fmt.Sprintf("%s.html", filepath.Base(fname))
	fmt.Println("New html file name", outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}
	return nil
}

// parseContent take []byte, pass the bytes to blackfriday for convert to html
// then pass the content to bluemonday for final html
func parseContent(input []byte) []byte {
	// Parse the markdown file through blackfriday and bluemonday
	// to generage a valid and safe html UGCPolicy compliant then save
	// the final content to body variable
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Create a bytes buffer to write to fileName
	var buffer bytes.Buffer

	// Write html to bytes Buffer
	buffer.WriteString(header) // Predefine html header
	buffer.Write(body)         // .md content to html's body
	buffer.WriteString(footer) // predefine html footer

	return buffer.Bytes() // return the []byte back to caller
}

// saveHTML takes the final html []byte and the html filename and
// os.WriteFile to the html file
func saveHTML(outFname string, data []byte) error {
	//Write the bytes to the fileName
	return os.WriteFile(outFname, data, 0644)
}
