// O(n) Time complexity 
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 2, 1, -3, 4, 3, 1, -8, 6, -2, -1}
	fmt.Println(max12(arr))

}
func max12(arr []int) int {
	hashmap := make(map[int]int)
	max_len := 0
	curr_sum := 0
	for i := 0; i < len(arr); i++ {
		curr_sum += arr[i]
		if curr_sum == 0 {
			max_len = i + 1
		}
		if val, ok := hashmap[curr_sum]; ok {
			max_len = int(math.Max(float64(max_len), float64(i-val)))
		} else {
			hashmap[curr_sum] = i
		}

	}
	return max_len
}
