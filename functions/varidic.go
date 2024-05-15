package main

import "fmt"

func main() {
	a := check(10, 20, 30)
	fmt.Println(a)
	fmt.Println(" i am inside golan")

}
func check(c ...int) []int {

	return c
}
