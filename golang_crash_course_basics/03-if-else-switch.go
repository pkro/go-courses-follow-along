package main

import "fmt"

func main() {
	var a = 4.0
	var b = 3.0

	// parenthesis are optional, convention is not to use them
	if a == b {
		fmt.Println("false")
	} else if a > b {
		fmt.Println("true")
	} else {
		fmt.Println("not happening")
	}

	// intitialization in the if statement
	if frac := a / b; frac > 0.5 {
		fmt.Printf("%v is more thatn half of %v\n", a, b)
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

	// naked switch
	switch {
	case a > b:
		fmt.Println("%v is greater than %v", a, b)

	case a < b:
		fmt.Println("%v is smaller than %v", a, b)

	default:
		fmt.Println("%v == %v", a, b)
	}
}
