package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func dirents(dir string) (entries []os.FileInfo) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du : %d\n", err)
		entries = nil
		return
	}
	return
}

func walkDir(dir string, n *sync.WaitGroup, sb chan<- int64) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, n, sb)
		} else {
			sb <- entry.Size()
		}
	}
}

func size(dir string) (s int64) {
	var wg sync.WaitGroup
	sizeBytes := make(chan int64)

	go func() {
		for sb := range sizeBytes {
			s += sb
		}
	}()

	wg.Add(1)
	walkDir(dir, &wg, sizeBytes)
	wg.Wait()

	close(sizeBytes)
	return
}

func main() {
	roots := os.Args[1:]

	for _, root := range roots {
		fmt.Printf("sizeof %s = %.1gGB\n", root, float64(size(root))/1e9)
	}
}
