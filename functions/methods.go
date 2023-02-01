package main

import "fmt"

type guru struct {
	first string
	last  string
}
type mahe struct {
	guru
	isreal bool
}

func main() {
	vis := guru{
		first: "vishnu",
		last:  "katta",
	}
	fmt.Println(vis)
	value := mahe{
		guru: guru{
			first: "gurunath",
			last:  "reddy",
		},
		isreal: true,
	}
	value1 := mahe{
		guru: guru{
			first: "mahesh",
			last:  "reddy",
		},
		isreal: true,
	}
	value.happy()
	value1.happy()
	vis.happy()

}
func (m mahe) happy() {
	fmt.Println("i am", m.first, m.last)

}
func (g guru) happy() {
	fmt.Println("i am", g.first, g.last)

}
