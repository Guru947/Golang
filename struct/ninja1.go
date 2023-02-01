package main

import "fmt"

type person struct {
	firstname        string
	lastname         string
	favoriteicecream []string
}

func main() {
	a := person{
		firstname:        "vishnu",
		lastname:         "katta",
		favoriteicecream: []string{"chocolate", "butterschock"},
	}
	m := map[string]person{
		a.lastname: a,
	}
	fmt.Println(m)
	fmt.Println(a)
	fmt.Println(a.firstname, a.lastname, a.favoriteicecream)
	for i, v := range a.favoriteicecream {
		fmt.Println(i, v)
	}
	for i, v := range m {
		fmt.Println(i, v)
	}

}
