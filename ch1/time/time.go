package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())

	fmt.Println(time.UnixDate)
	fmt.Println(time.Now().Format(time.UnixDate))

	for _, f := range os.Args[1:] {
		fmt.Println(time.Now().Format(f))
	}
}
