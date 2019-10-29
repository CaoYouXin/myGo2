package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		checkErr(err)

		fmt.Printf("HTTP Status: %s\n", resp.Status)
		// body, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		fmt.Println()
		checkErr(err)

		// fmt.Printf("%s\n", body)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
}
