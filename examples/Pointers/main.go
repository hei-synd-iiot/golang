package main

import "fmt"

func zeroval(ival int) { // (1)!
	ival = 0
}

func zeroptr(iptr *int) { // (2)!
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // (3)!
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // (4)!
}
