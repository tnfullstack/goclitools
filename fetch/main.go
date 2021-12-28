// Fetch print the contect found at the URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	httpPre = "http://"
	url     string
)

func main() {
	for _, url = range os.Args[1:] {

		if !strings.HasPrefix(url, httpPre) {
			url = httpPre + url
			fmt.Println(url)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
