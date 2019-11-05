package main

import "fmt"

// ByteCounter counts the len of writen bytes
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // need same type
	return len(p), nil
}

func main() {
	var c ByteCounter
	fmt.Fprint(&c, "hello")
	fmt.Println(c)

	c = 0 // reset
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
