package main

import "fmt"

func main() {
	a := "guru"
	b := &a
	fmt.Printf("the b is this %T and it has this %v and its Address is %p\n", b, b,&b)
	fmt.Printf("the address of a is %p\n", &a)
}
