package main

import "fmt"

type abc struct {
	name string
	age  int
}

func change(a abc) {
	a.name = "guru"
	a.age = 22

}
func changeusingpointer(a *abc) {
	a.name = "guru"
	a.age = 22

}

func main() {
	a := abc{
		name: "mahe",
		age:  22,
	}
	fmt.Println("before", a)
	change(a)
	fmt.Println("after", a)
	fmt.Println("before using a pointer function", a)
	changeusingpointer(&a)
	fmt.Println("after using a pointer function", a)

}
