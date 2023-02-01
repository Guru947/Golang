package main

import "fmt"

func main() {
	x := vishnu()
	fmt.Printf("%T", x)
	fmt.Println(x())
	fmt.Println(vishnu())

}
func vishnu() func() int {
	return func() int {
		return 124
	}
}
