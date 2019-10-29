package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	printResults(counts)
}

func printResults(counts map[string]int) {
	fmt.Println("\r \rHere is the results:")
	for line, n := range counts {
		if n > 1 && len(line) > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
