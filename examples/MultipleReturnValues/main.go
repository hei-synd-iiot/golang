package main

import "fmt"

func vals() (int, int) { // (1)!
	return 3, 7
}

func main() {
	a, b := vals() // (2)!
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // (3)!
	fmt.Println(c)
}
