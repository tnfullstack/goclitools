package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// html header
const (
	defaultTemplate = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{ .Title }}</title>
	</head>
	<body>
	{{ .Body }}
	</body>
</html>
`
)

// content type represents the HTML content to add into the template
type content struct {
	Title string
	Body  template.HTML
}

// program entry
func main() {
	// Parse flags
	fileName := flag.String("file", "", "Markdown file to repview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()

	// If user did not provide input file, show usage
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}
	// Call and pass the cmd flag and fileName to run func
	if err := run(*fileName, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run take fname from cli, io.Writer, and a flag
func run(fn, tFn string, out io.Writer, skp bool) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	// call and pass the .md []byte to parseContent func for html parsing
	htmlData, err := parseContent(input, tFn)
	if err != nil {
		return err
	}

	// Create temporary file and check for errors
	temp, err := ioutil.TempFile("/tmp", "mdp*.html")
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
	if skp {
		return nil
	}
	defer os.Remove(outName)
	return preview(outName)
}

// parseContent take []byte, pass the bytes to blackfriday for convert to html
// then pass the content to bluemonday for final html
func parseContent(input []byte, tFn string) ([]byte, error) {
	// Parse the markdown file through blackfriday and bluemonday
	// to generage a valid and safe html UGCPolicy compliant then save
	// the final content to body variable
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Parse the contents of the defaultTemplate const into a new template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	// If user provided alternate template file, replace the defaultTemplate
	if tFn != "" {
		t, err = template.ParseFiles(tFn)
		if err != nil {
			return nil, err
		}
	}

	// Instantiate the content type, adding the title and body
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	// Create a bytes buffer to write to fileName
	var buffer bytes.Buffer

	// Execute the template with the content type
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil // return the []byte back to caller
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
	err = exec.Command(cPath, cParams...).Run()

	// Give the browser some time to open file before deleting it
	time.Sleep(2 * time.Second)
	return err
}
