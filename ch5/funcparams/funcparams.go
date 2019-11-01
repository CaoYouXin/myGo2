package main

import "fmt"

func modArr1(a [1]int) {
	a[0] = 1
}

func modArr2(a *[1]int) {
	a[0] = 1
}

func modSlice1(s []int) {
	s[0] = 1
}

func modSlice2(s *[]int) {
	(*s)[0] = 1
}

func modSlice3(s []int) {
	s = append(s, 1)
}

func main() {
	arr1 := [1]int{0}
	modArr1(arr1)
	fmt.Println(arr1)

	arr2 := [1]int{0}
	modArr2(&arr2)
	fmt.Println(arr2)

	slice1 := []int{0}
	modSlice1(slice1)
	fmt.Println(slice1)

	slice2 := []int{0}
	modSlice2(&slice2)
	fmt.Println(slice2)

	slice3 := []int{0}
	modSlice3(slice3)
	fmt.Println(slice3)
}
