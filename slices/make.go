package main

import "fmt"

func main() {
	a := make([]int, 10, 11)
	fmt.Println(a, len(a), cap(a))
	a[0] = 1
	b := []int{15, 16}
	fmt.Println(a, len(a), cap(a))
	a = append(a, 12, 13)
	fmt.Println(a, len(a), cap(a))
	a = append(a, b...)
	fmt.Println(a, "\t", len(a))

}
