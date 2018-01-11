package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Sampler samples occurrences.
type Sampler struct {
	n       int
	skipped int
}

// NewSampler returns a Sampler that samples 1 in n occurrences.
func NewSampler(n int) Sampler {
	return Sampler{
		n: n,
	}
}

// Sample returns true if the occurrence should be sampled.
func (s *Sampler) Sample() bool {
	s.skipped++
	if s.skipped == s.n {
		s.skipped = 0
		return true
	}
	return false
}

func main() {
	n := flag.Uint("n", 1, "1 in n lines to sample. 1 means every line is printed to stdout.")
	flag.Parse()

	sampler := NewSampler(int(*n))
	reader := bufio.NewReader(os.Stdin)

	var line string
	var err error
	for line, err = reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		if sampler.Sample() {
			fmt.Print(line)
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting on error: %v", err)
		os.Exit(1)
	}
}
