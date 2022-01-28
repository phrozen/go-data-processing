package main

import (
	"bufio"
	"log"
	"os"

	"github.com/phrozen/processor/pkg/counter"
)

const BUFFER_SIZE = 64 * 1024 * 1024

func OpenFile(filename string) *os.File {
	log.Println("Opening file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	stats, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Size:", stats.Size())
	return file
}

func main() {
	// Open the file
	file := OpenFile(os.Args[1])
	// Close the file when we are done using it
	defer file.Close()
	// Create a new counter
	count := counter.NewCounter()
	// Create new Scanner from the file
	scanner := bufio.NewScanner(file)
	// Create a new buffer with a fixed size
	buffer := make([]byte, BUFFER_SIZE)
	// Assing buffer to our scanner
	scanner.Buffer(buffer, BUFFER_SIZE)
	// For as long as we can scan...
	for scanner.Scan() {
		// Process line
		count.ProcessLine(scanner.Text())
	}
	// Check the error of the scanner if any
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	// Print results
	count.PrintResult()
}
