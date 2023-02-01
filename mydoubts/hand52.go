package main

import (
	"fmt"
	"math"
)

type square struct {
	radius float64
}
type circle struct {
	length float64
}

func (s square) area() float64 {
	return math.Pi * s.radius * s.radius
}
func (c circle) area() float64 {
	return c.length * c.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	a := circle{
		length: 53.0,
	}
	b := square{
		radius: 14.555,
	}
	info(a)
	info(b)

}
