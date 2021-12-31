package main

import (
	"log"
	"os"
)

// init() is call prior to main
func init() {
	// change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

// main starting point of the program
func main() {

	// Perform a search
	search.Run("Presedent")

}
