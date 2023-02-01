package main

import "fmt"

func main() {
	a := 10
	b := &a
	c := &b
	fmt.Printf("%v\t%p\n", a, &a)
	fmt.Printf("%v\t%p\n", b, &b)
	fmt.Printf("%v\t%p\n", c, &c)
	fmt.Println(b)

}
