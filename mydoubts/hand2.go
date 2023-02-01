package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := guru1(a...)
	fmt.Println(result)
	ab := gur(a)
	fmt.Println(ab)

}
func guru1(a ...int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
func gur(a []int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
