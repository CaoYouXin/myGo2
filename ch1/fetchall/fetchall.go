package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	s := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(s, url)
		go fetch(s, url)
	}

	for range os.Args[1:] {
		fmt.Println(<-s)
		fmt.Println(<-s)
	}

	fmt.Printf("%.2fs elapsed.\n", time.Since(start).Seconds())
}

func fetch(s chan string, url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		s <- fmt.Sprintf("%v", err)
		return
	}

	bytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		s <- fmt.Sprint(err)
		return
	}

	elapsed := time.Since(start).Seconds()
	s <- fmt.Sprintf("%.2fs %dbytes %s", elapsed, bytes, url)
}
