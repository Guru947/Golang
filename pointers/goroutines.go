package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("main program starts execution ")
	var wg sync.WaitGroup
	wg.Add(1)

	go f1(&wg)
	wg.Wait()
	f2()

	fmt.Println("main program stops")

}
func f1(wg *sync.WaitGroup) {
	fmt.Println("the f1 function starts execution ")
	wg.Done()
	fmt.Println("the f1 function stops ")

}
func f2() {
	fmt.Println("the f2 function starts execution")

	fmt.Println("F2 stops")

}
