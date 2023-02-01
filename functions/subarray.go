//program with O(n**2) Time complexity
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 2, 1, -3, 4, 3, 1, -8, 6, -2, -1}
	fmt.Println("largest  subarray with length ", maxi(arr))

}
func maxi(arr []int) int {
	length := 0
	for i := 0; i < len(arr); i++ {
		var sum int
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum == 0 {
				length = int(math.Max(float64(length), float64(j-i+1)))
			}
		}
	}
	return length

}
