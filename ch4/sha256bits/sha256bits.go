package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	sha1 := sha256.Sum256([]byte(os.Args[1]))
	sha2 := sha256.Sum256([]byte(os.Args[2]))

	var count int64
	for i := 0; i < sha256.Size; i++ {
		if sha1[i] != sha2[i] {
			count++
		}
	}

	fmt.Printf("Output: %d\n\t%x\n\t%x\n", count, sha1, sha2)
}
