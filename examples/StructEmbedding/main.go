package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base // (1)!
	str  string
}

func main() {

	co := container{
		base: base{ // (2)!
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // (3)!

	fmt.Println("also num:", co.base.num) // (4)!

	fmt.Println("describe:", co.describe()) // (5)!

	type describer interface {
		describe() string
	}

	var d describer = co // (6)!
	fmt.Println("describer:", d.describe())
}
