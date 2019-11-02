package main

import "fmt"

func recoverMagic() (res int) {
	defer func() {
		switch p := recover(); p {
		case nil:
		case 0:
			res = 1
		default:
			panic(p)
		}
	}()
	panic(res)
}

func main() {
	fmt.Println(recoverMagic())
}
