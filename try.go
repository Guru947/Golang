package main

import "fmt"

func main() {

	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
	a[2] = 4
	fmt.Println(a)

}
