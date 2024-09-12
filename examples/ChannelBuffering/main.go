package main

import "fmt"

func main() {

	messages := make(chan string, 2) // (1)!

	messages <- "buffered" // (2)!
	messages <- "channel"

	fmt.Println(<-messages) // (3)!
	fmt.Println(<-messages)
}
