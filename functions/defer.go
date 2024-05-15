package main

import "fmt"

func main() {
	a()
	defer b()
	defer c()
	fmt.Println("I am inside vs code")

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
