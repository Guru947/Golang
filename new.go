package main

import "fmt"

type a struct {
	name string
	age  int
}
type bad struct {
	a
	girlfriend bool
}
type res interface {
	good()
}

func mam(r res) {
	switch r.(type) {
	case a:
		fmt.Println("iam in switch case of mam function", r.(a).name)
	case bad:
		fmt.Println("i am in switch case of man fucoion ", r.(bad).name)
	}
	fmt.Println("i am in mam fuction", r)
}

func (b bad) good() {
	fmt.Print(b.girlfriend, b.a)
}
func (a a) good() {
	fmt.Print(a.name)
}

func main() {
	value := bad{
		a: a{
			name: "guru",
			age:  22,
		},
		girlfriend: true,
	}
	value.good()

	value1 := a{
		name: "mahe",
		age:  22,
	}
	value1.good()
	mam(value)
	mam(value1)

}
