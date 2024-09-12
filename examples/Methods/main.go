package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // (1)!
	return r.width * r.height
}

func (r rect) perim() int { // (2)!
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area()) // (3)!
	fmt.Println("perim:", r.perim())

	rp := &r // (4)!
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
