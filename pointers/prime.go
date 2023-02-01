package main

import "fmt"

func main() {
	a := 10
	for i := 1; i < a; i++ {
		for j := 2; j < a-1; j++ {
			{
				if i%j == 0 {
					continue
				} else {
					fmt.Println(i)
				}

			}
		}
	}
}
