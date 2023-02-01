package main

import "fmt"

func main() {
	a()
	defer b()
	defer c()
}

func c() {
	fmt.Println("c")
}
func a() {
	fmt.Println("a")
}
func b() {
	fmt.Println("b")
}
