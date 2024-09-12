package main

import "fmt"

type person struct { // (1)!
	name string
	age  int
}

func newPerson(name string) *person { // (2)!
	p := person{name: name}
	p.age = 42
	return &p // (3)!
}

func main() {
	fmt.Println(person{"Bob", 20}) // (4)!

	fmt.Println(person{name: "Alice", age: 30}) // (5)!

	fmt.Println(person{name: "Fred"}) // (6)!

	fmt.Println(&person{name: "Ann", age: 40}) // (7)!

	fmt.Println(newPerson("Jon")) // (8)!

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // (9)!

	sp := &s
	fmt.Println(sp.age) // (10)!

	sp.age = 51 // (11)!
	fmt.Println(sp.age)

	dog := struct { // (12)!
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
