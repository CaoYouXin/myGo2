package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var t = flag.Int("t", 256, "specify crypto algrithm")

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		switch {
		case *t == 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(arg)))
		case *t == 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(arg)))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(arg)))
		}
	}
}
