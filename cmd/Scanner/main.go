package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Open file, it's an io.Reader
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Close it when we are done
	defer file.Close()
	// Create a new buffered scanner from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Process each line from the buffer
		ProcessLine(scanner.Text())
	}
	// Check the error of the scanner if any
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var count uint64

func ProcessLine(line string) {
	count++
}
