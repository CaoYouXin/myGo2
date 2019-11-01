package main

import (
	"fmt"
	visithtml "go-starter/ch5/visitHTML"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func hasNoElementChild(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			return false
		}
	}
	return true
}

func main() {
	root, err := visithtml.FetchAndParse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "beautify html: %v\n", err)
		os.Exit(1)
	}

	t := 0
	tmpl1 := "%s<%s>\n"
	tmpl2 := "%s</%s>\n"
	tmpl3 := "%s<%s />\n"
	visithtml.Visit(root, func(n *html.Node) {
		if n.Type == html.ElementNode {
			empties := make([]string, t+1)
			prefix := strings.Join(empties, "  ")

			tmpl := tmpl1
			if hasNoElementChild(n) {
				tmpl = tmpl3
			}

			fmt.Printf(tmpl, prefix, n.Data)
			t++
		}
	}, func(n *html.Node) {
		if n.Type == html.ElementNode {
			t--
			if !hasNoElementChild(n) {
				empties := make([]string, t+1)
				prefix := strings.Join(empties, "  ")
				fmt.Printf(tmpl2, prefix, n.Data)
			}
		}
	})
}
