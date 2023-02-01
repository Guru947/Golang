package main

import "fmt"

type a struct {
	first string
	age   int
}
type b struct {
	a
	last  string
	phnno int
}

func main() {
	result := b{
		a: a{
			first: "guru",
			age:   22,
		},
		last:  "reddy",
		phnno: 123,
	}
	fmt.Println(result)
	fmt.Println(result.last, result.first, result.age, result.phnno)

}
