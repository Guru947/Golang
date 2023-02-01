package main

import "fmt"

type gur struct {
	name string
	age  int
}
type mahe struct {
	phnn int
}

func (g gur) add() {
	fmt.Println(g.name, g.age)
}
func (m mahe) add1() {
	fmt.Println(m.phnn)
}

func main() {
	a := gur{
		name: "guru",
		age:  22,
	}
	//b := mahe{phnn: 124435}
	a.add()
	//b.add1()
	for i, v := range a.name {
		fmt.Println(i, string(v))
	}

}
