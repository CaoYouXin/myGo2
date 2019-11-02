package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	filename = path.Base(resp.Request.URL.Path)
	if filename == "/" {
		filename = "index.html"
	}

	f, err := os.Create(filename)
	if err != nil {
		return
	}

	defer func() {
		err = f.Close()
	}()

	n, err = io.Copy(f, resp.Body)

	return
}

func main() {
	fmt.Println(fetch(os.Args[1]))
}
