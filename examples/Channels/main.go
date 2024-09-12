package main

import "fmt"

func main() {

	messages := make(chan string) // (1)!

	go func() { messages <- "ping" }() // (2)!

	msg := <-messages // (3)!
	fmt.Println(msg)
}
