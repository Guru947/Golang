package main

import "fmt"

func main() {
	boo()
	bar()
	fmt.Println(boo())
	fmt.Println(bar())
}
func boo() int {
	return 45
}
func bar() (int, string) {
	return 87, "gurunath"
}
