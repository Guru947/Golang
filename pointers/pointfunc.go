package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	ab(&a)

}
func ab(c *int) {
	fmt.Println(*c)
	*c = 11
	fmt.Println(*c)

}
