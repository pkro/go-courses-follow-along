package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings must be defined with double quotes or backtick (raw strings)
	book := "Power of habit"

	// part of strings can be accessed with slices
	fmt.Println(book[6:])  // of habit
	fmt.Println(book[:6])  // Power
	fmt.Println(book[6:9]) // of

	// strings are immutable (but not re-assignable)
	// book[0] = 116 // error

	// strings are unicode enables
	book = "⚓"

	// raw strings ignore special characters (\n etc) and can be written on multiple lines
	poem := `Roses are red
some are blue`
	fmt.Println(poem)

	// split by whitespace in array
	fmt.Println(strings.Fields("a b    c")) // [a b c]
	// split by any character
	fmt.Println(strings.Split(strings.ToUpper("a:b:c"), ":")) // [A B C]

}
