package main

import (
	"fmt"
	"math"
)

type geometry interface { // (1)!
	area() float64
	perim() float64
}

type rect struct { // (2)!
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 { // (3)!
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 { // (4)!
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { // (5)!
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) // (6)!
	measure(c)
}
