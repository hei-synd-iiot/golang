package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string // (1)!
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // (2)!
	fmt.Println("em:", s, "le:", len(s), "ca:", cap(s))

	s[0] = "a" // (3)!
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // (4)!

	s = append(s, "d") // (5)!
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // (6)!
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] // (7)!
	fmt.Println("sl1:", l)

	l = s[:5] // (8)!
	fmt.Println("sl2:", l)

	l = s[2:] // (9)!
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // (10)!
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"} // (11)!
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) // (12)!
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
