package main

import "fmt"

type checkIsPalindrome string

func (c *checkIsPalindrome) Less(i, j int) bool {
	return (*c)[i] < (*c)[j]
}

func (c *checkIsPalindrome) IsPalindrome() bool {
	for i, j := 0, len(*c)-1; i < j; i, j = i+1, j-1 {
		if c.Less(i, j) != c.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	a, b := checkIsPalindrome("aba"), checkIsPalindrome("abc")
	fmt.Println(a.IsPalindrome(), b.IsPalindrome())
}
