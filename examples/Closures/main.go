package main

import "fmt"

func intSeq() func() int { // (1)!
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq() // (2)!

	fmt.Println(nextInt()) // (3)!
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // (4)!
	fmt.Println(newInts())
}
