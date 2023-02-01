package main

import "fmt"

func main() {
	age := 33
	name := "guru"
	single := true
	fmt.Println("before calling a  change function", age, name, single)
	change(age, name, single)
	fmt.Println("after calling a  change function", age, name, single)
	changeWithPointers(&age, &name, &single)
	fmt.Println("before calling a function", age, name, single)
}
func change(a int, n string, s bool) {
	a = 22
	n = "nath"
	s = true
}
func changeWithPointers(a *int, n *string, s *bool) {

	*a = 22
	*n = "nath"
	*s = true
}
