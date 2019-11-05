package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Counter for words and lines
type Counter struct {
	words, lines int
}

// ScanWords scans the string(p) returns words count
func (c *Counter) ScanWords(p []byte) (int, error) {
	r := strings.NewReader(string(p))
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	sum := 0
	for s.Scan() {
		sum++
	}

	return sum, nil
}

// ScanLines scans the string(p) returns lines count
func (c *Counter) ScanLines(p []byte) (int, error) {
	r := strings.NewReader(string(p))
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	sum := 0
	for s.Scan() {
		sum++
	}

	return sum, nil
}

func (c *Counter) Write(p []byte) (int, error) {
	words, err := c.ScanWords(p)
	if err == nil {
		c.words += words
	}

	lines, err := c.ScanLines(p)
	if err == nil {
		c.lines += lines
	}

	return 0, nil
}

func (c *Counter) String() string {
	return fmt.Sprintf("%d words, %d lines", c.words, c.lines)
}

func main() {
	var c Counter
	fmt.Fprint(&c, "Hello, World!")
	fmt.Println(&c)
}
