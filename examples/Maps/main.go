package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // (1)!

	m["k1"] = 7 // (2)!
	m["k2"] = 13

	fmt.Println("map:", m) // (3)!

	v1 := m["k1"] // (4)!
	fmt.Println("v1:", v1)

	v3 := m["k3"] // (5)!
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m)) // (6)!

	delete(m, "k2") // (7)!
	fmt.Println("map:", m)

	clear(m) // (8)!
	fmt.Println("map:", m)

	_, ok := m["k2"] // (9)!
	fmt.Println("prs:", ok)

	n := map[string]int{"foo": 1, "bar": 2} // (10)!
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) { // (11)!
		fmt.Println("n == n2")
	}
}
