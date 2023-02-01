package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println(a, b)
	ab(a, b)
	fmt.Println(a, b, "after function calling")

}
func ab(a, b int) {
	a, b = 20, 30
	fmt.Println(a, b, "i am inside the function")
}
