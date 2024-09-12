package main

import (
	"embed" // (1)!
)

//go:embed resources/single_file.txt
var fileString string // (2)!

//go:embed resources/single_file.txt
var fileByte []byte // (3)!

//go:embed resources/single_file.txt
//go:embed folder/*.hash
var folder embed.FS // (4)!

func main() {

	println(fileString) // (5)!
	println(string(fileByte))

	content1, _ := folder.ReadFile("resources/file1.hash") // (6)!
	println(string(content1))

	content2, _ := folder.ReadFile("resources/file2.hash")
	println(string(content2))
}
