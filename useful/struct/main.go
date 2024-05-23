package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	a := person{}
	a.name = "guru"
	a.age = 30
	fmt.Println(a)
	result := check("mahe", 29)
	fmt.Println(*(result))
}

//done some operatin using function

func check(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
