package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	var buf bytes.Buffer

	for i := len(s) % 3; i <= len(s); i += 3 {
		if i == 0 {
			continue
		}

		if i < 3 {
			buf.WriteString(s[0:i])
		} else {
			buf.WriteString(s[i-3 : i])
		}

		if i != len(s) {
			buf.WriteByte(',')
		}
	}

	fmt.Println(buf.String())
}
