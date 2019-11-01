package visithtml

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Visit visit html doc
func Visit(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Visit(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// FetchAndParse fetch & parse to root html node
func FetchAndParse(url string) (root *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}

	root, err = html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return
}

// ElementByID find element node by id
func ElementByID(n *html.Node, id string) *html.Node {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return n
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		found := ElementByID(c, id)
		if found != nil {
			return found
		}
	}

	return nil
}
