package main

import "fmt"

func main() {
	arr := []int{2, 2, 1, -3, 4, 3, 1, -8, 6, -2, -1}
	local(arr)

}
func local(arr []int) {

	for i := 1; i < len(arr)-2; i++ {
		if arr[0] < arr[1] {
			fmt.Println("the local minima is at index ", 0)
			break
		}
		if (arr[i] < arr[i-1]) && (arr[i] < arr[i+1]) {
			fmt.Println("the local minima is at index ", i)
			break

		}
		if arr[len(arr)-1] < arr[len(arr)-2] {
			fmt.Println("the local minima is at index", (len(arr) - 1))
			break
		}

	}

}
