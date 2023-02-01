package main

import "fmt"

func main() {
	a := recursion(5)
	fmt.Println(a)

}
func recursion(a int) int {
	if a == 0 {
		return 1
	}
	return a * recursion(a-1)
}
