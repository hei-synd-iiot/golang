package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี" // (1)!

	fmt.Println("Len:", len(s)) // (2)!

	for i := 0; i < len(s); i++ { // (3)!
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // (4)!

	for idx, runeValue := range s { // (5)!
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:]) // (6)!
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue) // (7)!
	}
}

func examineRune(r rune) {
	if r == 't' { // (8)!
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
