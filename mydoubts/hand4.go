package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	a := person{
		first: "guru", last: "nath", age: 22,
	}
	b := person{
		first: "vishnu",
		last:  "katta",
		age:   45,
	}
	a.speak()
	b.speak()
}
func (p person) speak() {
	fmt.Println("my name is :", p.first, "and my age", p.age)
}
