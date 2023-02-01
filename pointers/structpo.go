package main

import "fmt"

type person struct {
	name string
	age  int
}

func change(p person) {
	p.name = "mahessh"
	p.age = 33
}
func changebypointer(p *person) {
	p.name = "mahessh"
	p.age = 33
}

func main() {
	a := person{
		name: "guru",
		age:  22,
	}
	fmt.Println("before calling the function ", a)
	change(a)
	fmt.Println("after calling the function ", a)
	fmt.Println("before calling the function of pointers ", a)
	changebypointer(&a)
	fmt.Println("after calling the function of pointers ", a)

}
