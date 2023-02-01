package main

import "fmt"

func main() {

	fmt.Println(guru()())

}
func guru() func() int {
	return func() int {
		return 123
	}
}
