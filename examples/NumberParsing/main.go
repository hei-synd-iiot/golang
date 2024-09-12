package main

import (
	"fmt"
	"strconv" // (1)!
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64) // (2)!
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) // (3)!
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) // (4)!
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) // (5)!
	fmt.Println(u)

	k, _ := strconv.Atoi("135") // (6)!
	fmt.Println(k)

	_, e := strconv.Atoi("wat") // (7)!
	fmt.Println(e)
}
