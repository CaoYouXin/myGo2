package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout

	if _, ok := w.(*os.File); ok {
		fmt.Println(w)
	}

	fmt.Println(w.(io.Writer))

	b, ok := w.(*bytes.Buffer)
	fmt.Println(b, ok)

	_, err := os.Open("/no/such/file")
	if os.IsNotExist(err) {
		fmt.Println(err)
	}
}
