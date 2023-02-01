package main

import "fmt"

func main() {

	a := []int{
		1, 2, 3, 4, 5, 6,
	}

	b := []int{
		12, 34, 5, 6, 7, 7, 8,
	}

	c := [][]int{a, b}
	fmt.Println(c)

}
