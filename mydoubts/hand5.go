package main

import "fmt"

type square struct {
	area int
}
type circle struct {
	area int
}
type shape interface {
	AREA() int
}

func info(s shape) {
	fmt.Println(s.AREA())

}
func (c circle) AREA() int {
	return c.area

}
func (s square) AREA() int {
	return s.area

}

func main() {
	a := square{
		area: 120,
	}
	b := circle{
		area: 130,
	}
	info(a)
	info(b)

}
