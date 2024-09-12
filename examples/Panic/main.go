package main

import "os"

func main() {

	panic("a problem") // (1)!

	_, err := os.Create("/tmp/file") // (2)!
	if err != nil {
		panic(err)
	}
}
