package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	idx := strings.Index(s, "foo")
	if idx < 0 {
		return s
	}
	return s[0:idx] + f("foo") + expand(s[idx+3:], f)
}

func main() {
	fmt.Println(expand("foobarfoo", func(foo string) string { return "bar" }))
}
