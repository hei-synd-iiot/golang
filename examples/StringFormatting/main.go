package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p) // (1)!

	fmt.Printf("struct2: %+v\n", p) // (2)!

	fmt.Printf("struct3: %#v\n", p) // (3)!

	fmt.Printf("type: %T\n", p) // (4)!

	fmt.Printf("bool: %t\n", true) // (5)!

	fmt.Printf("int: %d\n", 123) // (6)!

	fmt.Printf("bin: %b\n", 14) // (7)!

	fmt.Printf("char: %c\n", 33) // (8)!

	fmt.Printf("hex: %x\n", 456) // (9)!

	fmt.Printf("float1: %f\n", 78.9) // (10)!

	fmt.Printf("float2: %e\n", 123400000.0) // (11)!
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "string") // (12)!

	fmt.Printf("str2: %q\n", "string") // (13)!

	fmt.Printf("str3: %x\n", "hex this") // (14)!

	fmt.Printf("pointer: %p\n", &p) // (15)!

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // (16)!

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) // (17)!

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // (18)!

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // (19)!

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b") // (20)!

	s := fmt.Sprintf("sprintf: a %s", "string") // (21)!
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // (22)!
}
