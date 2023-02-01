package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheels bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 5,
			color: "green",
		},
		fourWheels: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 6,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(t, s)
	fmt.Println(t.color, t.doors, t.fourWheels, t.vehicle)
	fmt.Println(s.color, s.doors, s.luxury, s.vehicle)

}
