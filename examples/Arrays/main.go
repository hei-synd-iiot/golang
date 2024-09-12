package main

import "fmt"

func main() {
	var a [5]int // (1)!
	fmt.Println("emp:", a)

	a[4] = 100 // (2)!
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // (3)!

	b := [5]int{1, 2, 3, 4, 5} // (4)!
	fmt.Println("dcl:", b)

	var twoD [2][3]int // (5)!
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
