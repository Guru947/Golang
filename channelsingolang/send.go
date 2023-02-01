package main

import "fmt"

func main() {
	c := make(chan int)
	go guru(c)
	mahe(c)
}
func guru(c chan<- int) {

	c <- 42
}
func mahe(c <-chan int) {

	fmt.Println("this is recieving channels", <-c)
}
