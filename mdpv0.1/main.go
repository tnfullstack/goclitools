package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

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
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()

	// If user did not provide input file, show usage
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}
	// Call and pass the cmd flag and fileName to run func
	if err := run(*fileName, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fname string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	// call and pass the .md []byte to parseContent func for html parsing
	htmlData := parseContent(input)

	// Create temporary file and check for errors
	temp, err := ioutil.TempFile("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Println(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}
	return preview(outName)
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

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	// Define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	// Append filename to parameters slice
	cParams = append(cParams, fname)

	// Locate executable in PATH
	cPath, err := exec.LookPath(cName)

	if err != nil {
		return err
	}

	// Open the file using default program
	return exec.Command(cPath, cParams...).Run()
}
