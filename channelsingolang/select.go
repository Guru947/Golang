package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go guru(even, odd, quit)
	mahe(even, odd, quit)
	fmt.Println("this is going to qiut")

}
func guru(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0

}
func mahe(e, o, q <-chan int) {
	for {
		select {
		case a := <-e:
			fmt.Println("this is from the even channel", a)
		case a := <-o:
			fmt.Println("this is from the odd channel", a)
		case a := <-q:
			fmt.Println("this is from the odd channel", a)

			return
		}
	}
}
