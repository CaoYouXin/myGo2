package main

import (
	"fmt"
	"go-starter/ch5/visitHTML"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// URL html parser
type URL string

// Parse parses a url string to *html.Node
func (url *URL) Parse() (*html.Node, error) {
	resp, err := http.Get(string(*url))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("new reader: %v", resp.Status)
	}

	node, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return node, nil
}

func main() {
	var url URL = "https://www.baidu.com"
	node, err := url.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "new reader: %v\n", err)
		os.Exit(1)
	}

	visithtml.Visit(node, func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Println(n.Data)
		}
	}, nil)
}
