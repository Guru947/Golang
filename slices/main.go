package main

import "fmt"

func main() {
	x := []string{
		"guru", "mahe",
	}
	for i, v := range x {
		fmt.Println(i, v)
		

	}
}
