package main

import (
	"bufio"
	"log"
	"os"

	"github.com/phrozen/go-data-processing/pkg/counter"
)

func main() {
	// Open file, it's an io.Reader
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Close it when we are done
	defer file.Close()
	// Create counter
	count := counter.NewCounter()
	// Create a new buffered scanner from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Process each line from the buffer
		count.ProcessLine(scanner.Text())
	}
	// Check the error of the scanner if any
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Results
	count.PrintResult()
}
