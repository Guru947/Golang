package main

import "fmt"

func main() {
	guru := 21

	fmt.Println("before changing", guru)
	change(guru)
	fmt.Println("after changing", guru)
	changewithpointer(&guru)
	fmt.Println(guru)

}
func change(g int) {
	g = 43
}
func changewithpointer(g *int) {
	*g = 45
}
