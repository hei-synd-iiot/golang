package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs) // (1)!
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints) // (2)!
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints) // (3)!
	fmt.Println("Sorted: ", s)
}
