package main

import "fmt"

func sum(nums ...int) { // (1)!
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums { // (2)!
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2) // (3)!
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // (4)!
}
