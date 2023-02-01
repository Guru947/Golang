package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Print("name\t")
	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}
	fmt.Print("loves\tor\thates\t")
	_, err = fmt.Scan(&b)
	if err != nil {
		panic(err)
	}
	fmt.Print("name\t")
	_, err = fmt.Scan(&c)
	if err != nil {
		panic(err)
	}
	fmt.Println(a, b, c)

}
