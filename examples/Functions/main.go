package main

import "fmt"

func plus(a int, b int) int { // (1)!

	return a + b // (2)!
}

func plusPlus(a, b, c int) int { // (3)!
	return a + b + c
}

func main() {
	res := plus(1, 2) // (4)!
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	p := plus // (5)!
	fmt.Println(p(1, 2))
}

func init() { // (6)!
	fmt.Println("init() called")
}
