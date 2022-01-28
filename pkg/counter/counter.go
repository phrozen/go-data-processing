package counter

import (
	"bytes"
	"log"
	"strings"
	"sync/atomic"
	"time"
)

type Counter struct {
	one   uint64
	two   uint64
	three uint64
	four  uint64
	more  uint64
	total uint64
	bytes uint64
	start time.Time
}

func NewCounter() *Counter {
	return &Counter{
		start: time.Now(),
	}
}

func (counter *Counter) ProcessLine(line string) {
	counter.total++
	counter.bytes += uint64(len(line))
	s := strings.Split(line, ":")
	switch s[1] {
	case "1":
		counter.one++
	case "2":
		counter.two++
	case "3":
		counter.three++
	case "4":
		counter.four++
	default:
		counter.more++
	}
}

func (counter *Counter) PrintResult() {
	elapsed := time.Since(counter.start)
	log.Println("Elapsed:", elapsed)
	t := float64(counter.bytes) / float64(elapsed) * 1000
	log.Printf("Throughput: %.2f MB/s\n", t)
	log.Printf("%+v\n", counter)
}

type AtomicCounter Counter

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{
		start: time.Now(),
	}
}

func (counter *AtomicCounter) ProcessLine(line []byte) {
	atomic.AddUint64(&counter.total, 1)
	atomic.AddUint64(&counter.bytes, uint64(len(line)))
	s := bytes.Split(line, []byte(":"))
	switch string(s[1]) {
	case "1":
		atomic.AddUint64(&counter.one, 1)
	case "2":
		atomic.AddUint64(&counter.two, 1)
	case "3":
		atomic.AddUint64(&counter.three, 1)
	case "4":
		atomic.AddUint64(&counter.four, 1)
	default:
		atomic.AddUint64(&counter.more, 1)
	}
}

func (counter *AtomicCounter) PrintResult() {
	elapsed := time.Since(counter.start)
	log.Println("Elapsed:", elapsed)
	t := float64(counter.bytes) / float64(elapsed) * 1000
	log.Printf("Throughput: %.2f MB/s\n", t)
	log.Printf("%+v\n", counter)
}
