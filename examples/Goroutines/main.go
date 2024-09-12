package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") // (1)!

	go f("goroutine") // (2)!

	go func(msg string) { // (3)!
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) // (4)!
	fmt.Println("done")
}
