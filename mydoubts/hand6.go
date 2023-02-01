package main

import "fmt"

func main() {
	func(a int) {
		result := 0
		for i := 0; i < 9; i++ {
			result += i
		}
		fmt.Println(result)

	}(2244)

}
