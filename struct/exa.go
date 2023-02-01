package main

import "fmt"

func main() {
	a := map[string]int{
		"guru": 22,
		"mahe": 20,
	}
	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i, v)
	}
}
