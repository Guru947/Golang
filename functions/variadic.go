package main

import "fmt"

func main() {
	xi :=
		[]int{
			1, 2, 3, 4,
		}
	result := guru(xi...)
	fmt.Println(result)

}
func guru(s ...int) int {
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	sum := 0
	for i, v := range s {
		fmt.Println("in the range ", i, "the value is to the index is ", v)
		sum += v
	}
	return sum

}
