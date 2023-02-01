package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println(a, b)
	ref(&a, &b)
	fmt.Println(a, b, "after function calling")

}
func ref(a, b *int) {
	*a, *b = 20, 30
	fmt.Println(*a, *b, "we are inside")
}
