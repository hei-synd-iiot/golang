package main

import (
	"fmt"
	"math"
)

const s string = "constant" // (1)!

const ( // (2)!
	c1 = 7
	c2 = 42
)

func main() {
	fmt.Println(s)

	const n = 500000000 // (3)!

	const d = 3e20 / n // (4)!
	fmt.Println(d)

	fmt.Println(int64(d)) // (5)!

	fmt.Println(math.Sin(n)) // (6)!
}
