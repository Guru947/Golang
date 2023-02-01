package main

import "fmt"

type person struct {
	address string
}

func main() {
	p := person{
		address: "sasanakota",
	}
	fmt.Println(p)
	changeme(&p)
	fmt.Println(p)
	fmt.Println(&p)

}
func changeme(c *person) {
	(*c).address = "parigi"
}
