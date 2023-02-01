package main

import "fmt"

func main() {

	a := struct {
		door int
		b    map[string]int
		c    []string
	}{

		door: 4,
		b: map[string]int{
			"guru": 22,
		},
		c: []string{
			"guru", "mahe",
		},
	}
	for i, v := range a.b {
		fmt.Println(i, v)
	}
	fmt.Println(a)
}
