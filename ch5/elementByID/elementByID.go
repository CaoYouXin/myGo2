package main

import (
	"fmt"
	visithtml "go-starter/ch5/visitHTML"
	"os"
)

func main() {
	root, err := visithtml.FetchAndParse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "beautify html: %v\n", err)
		os.Exit(1)
	}

	found := visithtml.ElementByID(root, os.Args[2])
	if found != nil {
		fmt.Println(found)
	}
}
