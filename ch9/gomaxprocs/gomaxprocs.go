package main

import "fmt"

// 不要运行
func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
