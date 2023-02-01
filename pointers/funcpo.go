package main

import "fmt"

func change(a *int) *float64 {
	*a = 22
	b := 5.55
	return &b
}
func changewithoutpointer(a *int) {
	*a = 55
}
func main() {
	a := 12
	b := &a
	fmt.Println("before calling a function", a)

	res := change(b)
	fmt.Println(res)
	fmt.Println("after calling a function", a)
	fmt.Println("before calling a function changewithout pointer", a)
	changewithoutpointer(b)
	fmt.Println("after calling a function change with out pointer", a)

}
