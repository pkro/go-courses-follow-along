package main

import "fmt"

func main() {
	myString := doStuff("hey")
	fmt.Println(myString)
}

func doStuff(what string) string {
	fmt.Println(what)
	return "stuff done"
}
