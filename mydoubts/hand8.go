package main

import "fmt"

func main() {
	value := []int{1, 2, 3, 4, 5, 6}
	res := even(value...)
	fmt.Println(res)
	res1 := alleven(even, value)
	fmt.Println(res1)

}
func even(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v

	}
	return sum
}
func alleven(f func(a ...int) int, va ...int) {
	var y []int
	for _, v := range va {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return (f(y...))

}
