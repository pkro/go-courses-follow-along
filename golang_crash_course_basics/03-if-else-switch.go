package main

import "fmt"

func main() {
	var a = 4
	var b = 3

	// parenthesis are optional, convention is not to use them
	if a == b {
		fmt.Println("false")
	} else if a > b {
		fmt.Println("true")
	} else {
		fmt.Println("not happening")
	}

	// no break necessary
	switch a {
	case 3:
		fmt.Println("it's 3")
	case 4:
		fmt.Println("it's 4")
	default:
		fmt.Println("it's probably not serious")
	}
}
