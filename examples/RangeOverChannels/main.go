package main

import "fmt"

func main() {

	queue := make(chan string, 2) // (1)!
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue { // (2)!
		fmt.Println(elem)
	}
}
