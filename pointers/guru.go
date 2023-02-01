package main

type person []int

func (p person) sum() {
	for _, v := range p {
		if v%2 == 0 {
			print(v, "\t")
		}
	}
}

func main() {
	a := person{1, 2, 3, 4, 5, 6}
	a.sum()

}
