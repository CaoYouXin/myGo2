package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	y, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Printf("GCD : %v\n", err)
	}

	for y != 0 {
		x, y = y, x%y
	}

	fmt.Printf("GCD result : %d\n", x)
}
