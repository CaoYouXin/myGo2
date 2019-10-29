package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s, time.Since(t))
}
