package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s, time.Since(t))
}
