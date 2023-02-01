package main

import "fmt"

type author struct {
	name string
	age  int
}

func (r author) search() {
	fmt.Println(r.name)
	fmt.Println(r.age)
}

func main() {
	a := author{name: "guru",
		age: 22,
	}
	a.search()

}
