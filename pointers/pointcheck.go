package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("%p\n", &a)
	fmt.Println(b)
	var ab *int
	fmt.Println(ab)
}
