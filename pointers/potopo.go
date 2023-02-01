package main

import "fmt"

func main() {

	a := 23
	b := &a
	c := &b
	fmt.Printf("the value of b %v,and the address of b %p\n ", b, &b)
	fmt.Printf("the value of b %p,and the address of b %v\n ", &b, b)
	fmt.Printf("the value of c %v,and the address of c %p\n ", c, &c)
	fmt.Printf("the value in b is os %v\n", b)
	fmt.Printf("the value in c is os %v\n", c)
	fmt.Printf("the value in *b is os %v\n", *b)
	fmt.Printf("the value in *c is os %v\n", *c)
}
