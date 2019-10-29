package main

import (
	"fmt"
	"go-starter/ch2/popcount"
	"os"
	"strconv"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Pop Count : %v\n", err)
		os.Exit(1)
	}

	fmt.Println(popcount.PopCount4(uint64(x)))
}
