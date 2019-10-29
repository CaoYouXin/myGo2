// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

func interruptted(c chan os.Signal, counts map[string]int) {
	<-c // c receives a ^C, handle it
	fmt.Println("\r \rHere is the results:")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	os.Exit(0)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	counts := make(map[string]int)

	go interruptted(c, counts)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Note: ignoring potential errors from input.Err()
}
