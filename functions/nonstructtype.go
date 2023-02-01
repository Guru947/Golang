package main

import "fmt"

type add1 int

func (a add1) addtion(b add1) add1 {
	return a * b

}

func main() {
	a := add1(20)
	b := add1(30)
	fmt.Println(a.addtion(b))

}
