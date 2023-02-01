package main

import "fmt"

type guru struct {
	first string
	last  string
}
type mahe struct {
	guru
	phn int
}
type person interface {
	c()
}

func bar(p person) {
	switch p.(type) {
	case guru:
		fmt.Println("it me type function in interfaces in case", p.(guru).first)
	case mahe:
		fmt.Println("it me type function in interface in case", p.(mahe).first)
	}
	fmt.Println(p, "it is interfaces")
}

func (m mahe) c() {
	fmt.Println("its me", m.first, m.last, m.phn)

}
func (g guru) c() {
	fmt.Println("its me", g.first, g.last)

}

func main() {
	a := guru{
		first: "gurunath",
		last:  "reddy",
	}
	b := mahe{
		guru: guru{
			first: "mahe",
			last:  "reddy",
		},
		phn: 123,
	}
	fmt.Println(a)
	fmt.Println(b)
	b.c()
	a.c()
	bar(a)
	bar(b)
}
