package main

import "fmt"

func main() {
	a := make([]int, 11, 10)
	a[10] = 19
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}
