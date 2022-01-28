package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Open, read and close the whole file
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Cast data to string as its a []byte
	// and split each line into an array of strings
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		// Process each line while the whole file
		// is in memory, file is already closed.
		ProcessLine(line)
	}
}

var count uint64

func ProcessLine(line string) {
	count++
}
