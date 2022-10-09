package main

import (
	"fmt"
	"os"
)

// return multiple values (return types MUST be in parentheses)
// it is idiomatic (but not syntactically required) to return an
// error as the last return value; the other returned values
// should return the default values for their type if there IS
// an error
func divmod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division by zero not allowed")
	}
	return a / b, a % b, nil
}

// If we encounter errors that can't be recovered from,
// we can use panic(error) to halt the program completely
func makeFileWeAbsolutelyNeed() {
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

func main() {
	// anonymous function with immediately invoked function expression (IIFE)
	func() {
		fmt.Println("hey ho")
	}()

	// no nested function declarations
	// func ohNo() {} // doesn't work

	// nested lambdas are OK
	ohYes := func() string { return "yes" }
	fmt.Println(ohYes())

	// receive multiple values from a function
	a, b, error := divmod(3, 4)

	if error == nil {
		fmt.Printf("%d - %d", a, b)
	} else {
		fmt.Println(error)
	}
}

// functions can be defined wherever
func doStuff(what string) string {
	fmt.Println(what)
	return "stuff done"
}
