package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(s[0])
	for _, c := range s {
		fmt.Printf("%q\t", c)
	}
	fmt.Println()
	r := []rune(s)
	fmt.Printf("%q\n", r[8])
	fmt.Printf("%t\n", strings.Contains(s, "世界"))
}
