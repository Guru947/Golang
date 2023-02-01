package main

import (
	"log"
	"os"
)

func main() {
	a()

}
func a() {
	_, err := os.Open("jdsfnh")
	if err != nil {
		log.Fatal(err)

	}

}
