package main

import (
	"bytes"
	"io"
)

func main() {
	var ioWriter io.Writer
	f(ioWriter)

	var buf *bytes.Buffer
	if buf != nil {
		f(buf)
	}
}

func f(out io.Writer) {
	if out == nil {
		return
	}
	out.Write([]byte("done!\n"))
}
