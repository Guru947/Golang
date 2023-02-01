package main

import "fmt"

func main() {
	a()
	defer b()
	c()

}
func a() {
	fmt.Println(
		"i am vishnu,",
	)
}
func b() {
	fmt.Println("i am waste fellow")
}
func c() {
	fmt.Println("i am lover boy  after becoming a lover boy i become a waste fellow")
}
