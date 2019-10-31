package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, i int) *tree {
	if t == nil {
		t = new(tree)
		t.value = i
		return t
	}

	if i < t.value {
		t.left = add(t.left, i)
	} else {
		t.right = add(t.right, i)
	}

	return t
}

func printResults(s []int, t *tree) []int {
	if t != nil {
		s = printResults(s, t.left)
		s = append(s, t.value)
		s = printResults(s, t.right)
	}
	return s
}

func sort(s []int) {
	var root *tree
	for _, i := range s {
		root = add(root, i)
	}
	printResults(s[:0], root)
}

func main() {
	var slice = []int{1, 3, 2}
	sort(slice)
	fmt.Println(slice)
}
