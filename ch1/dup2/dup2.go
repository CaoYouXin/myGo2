// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

func printResults(counts map[string]int) {
	fmt.Println("\r \rHere is the results:")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func interruptted(c chan os.Signal, counts map[string]int) {
	<-c // c receives a ^C, handle it
	printResults(counts)
	os.Exit(0)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Note: ignoring potential errors from input.Err()
}

func main() {
	counts := make(map[string]int)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go interruptted(c, counts)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		printResults(counts)
	}
}
