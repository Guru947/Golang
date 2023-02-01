package main

import "fmt"

func main() {
	vis := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(vis...))

}
func sum(a ...int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}
