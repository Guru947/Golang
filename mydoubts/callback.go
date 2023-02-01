package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := add(a1...)
	fmt.Println(res)
	RE1 := even(add, a1...)
	fmt.Println(RE1, "result of even numbers")
	RE2 := odd(add, a1...)
	fmt.Println(RE2, "result of odd numbers ")

}
func add(s1 ...int) int {

	result := 0
	for _, v := range s1 {

		result += v

	}
	return result
}
func even(f func(s1 ...int) int, av ...int) int {
	var y []int
	for _, v := range av {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return (f(y...))

}
func odd(f func(s1 ...int) int, va ...int) int {
	var y []int
	for _, v := range va {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return (f(y...))
}
