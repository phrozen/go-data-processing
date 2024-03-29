package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"sync"

	"github.com/phrozen/go-data-processing/pkg/counter"
)

const BUFFER_SIZE = 1024 * 1024 // 1MB per buffer

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
	defer file.Close()
	// Create pool for chunks
	pool := sync.Pool{
		New: func() interface{} {
			buf := make([]byte, BUFFER_SIZE)
			return buf
		},
	}
	// New atomic counter
	count := counter.NewAtomicCounter()
	// Wait group to sync go routines
	var wg sync.WaitGroup
	// Create new reader from the file
	r := bufio.NewReader(file)
	for {
		// Get buffer from pool
		buffer := pool.Get().([]byte)
		// Read at most len(buffer) from file
		n, err := r.Read(buffer)
		// At EOF we just break else panic
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		buffer = (buffer)[:n]
		// Read until next line break
		next, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		buffer = append(buffer, next...)
		// Write the chunk to the buffered channel for processing
		wg.Add(1)
		go func() {
			lines := bytes.Split((buffer)[0:len(buffer)-2], []byte("\r\n"))
			for _, line := range lines {
				count.ProcessLine(line)
			}
			pool.Put(buffer)
			wg.Done()
		}()
	}
	wg.Wait()
	count.PrintResult()
}
