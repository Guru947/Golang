package main

import "fmt"

func main() {
	a := []int{
		1, 3, 4, 5, 6, 7, 8,
	}

	a = append(a, 12, 454, 6, 6)
	fmt.Println(a)
	b := []int{
		12, 3, 4, 56,
	}
	a = append(a, b...)
	fmt.Println(a)
	a = append(a[:2], a[4:]...)
	fmt.Println(a)

}
