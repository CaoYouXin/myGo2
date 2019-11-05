package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// LReader is a limited reader
type LReader struct {
	read, limit int64
	origin      io.Reader
}

func (r *LReader) Read(p []byte) (n int, err error) {
	left := r.limit - r.read
	if left < int64(len(p)) {
		p = p[:left]
	}

	n, err = r.origin.Read(p)
	r.read += int64(n)

	if r.read == r.limit {
		err = io.EOF
		return
	}

	return
}

// LimitReader forge a limit reader
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LReader{read: 0, limit: n, origin: r}
}

func main() {
	r := strings.NewReader("Hello, WORLD!")
	lr := LimitReader(r, 5)
	s := bufio.NewScanner(lr)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
