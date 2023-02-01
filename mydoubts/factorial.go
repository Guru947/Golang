package main

import "fmt"

func main() {
	a := factorial(5)
	fmt.Println(a)

}
func factorial(a int) int {
	result := 1

	for ; a > 0; a-- {
		result *= a

	}
	return result

}
