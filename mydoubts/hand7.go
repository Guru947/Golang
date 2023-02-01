package main

import "fmt"

func main() {
	ab := func() {
		fmt.Println("i am happy")

	}
	ab()
	fmt.Printf("%T\n", ab)
}
