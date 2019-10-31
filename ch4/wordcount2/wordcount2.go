package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var counts = make(map[string]int)

type sortableCounts []string

func (sc sortableCounts) Len() int {
	return len(sc)
}

func (sc sortableCounts) Less(i, j int) bool {
	return counts[sc[i]] > counts[sc[j]]
}

func (sc sortableCounts) Swap(i, j int) {
	sc[i], sc[j] = sc[j], sc[i]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}

	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	sort.Sort(sortableCounts(words))
	fmt.Println("word\tcount:")
	for _, w := range words {
		fmt.Printf("%s\t%d\n", w, counts[w])
	}
}
