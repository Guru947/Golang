package main

import "fmt"

type guru struct {
	firstName  string
	secondName string
	age        int
}

func main() {
	a := guru{
		firstName:  "gurunath",
		secondName: "reddy",
		age:        22,
	}
	b := guru{
		firstName:  "mahesh",
		secondName: "reddy",
		age:        43,
	}
	fmt.Println(a, b)
	fmt.Println(a.firstName, a.secondName, a.age)
	fmt.Println(b.firstName, b.secondName, b.age)

}
