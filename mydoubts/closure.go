package main

import "fmt"

func main() {
	a := guru()
	b := guru()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}
func guru() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
