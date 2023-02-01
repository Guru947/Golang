package main

import "fmt"

func main() {
	fmt.Printf("%d", add(10, 20))
}
func add(a, b int) int {
	sum := a + b
	return sum
}
