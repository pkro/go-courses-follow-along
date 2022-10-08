package main

import "fmt"

func main() {
	myString := doStuff("hey")
	fmt.Println(myString)
	// anonymous function with immediately invoked function expression (IIFE)
	func() {
		fmt.Println("hey ho")
	}()
}

func doStuff(what string) string {
	fmt.Println(what)
	return "stuff done"

}
