package main

import (
	"fmt"
)

/*
Main datatypes:
string
bool
int (depending on system 32 or 64 signed integer)
int8 int16 int32 int64
uint8 uint16 uint32 uint64
uintptr (pointer to unsigned int)
byte (alias for uint8)
rune (alias for int32)
float32 float64
complex64 complex128
*/

// error - this initialization can only be used inside a function
// ohNo := "noooo"

// this is fine
var ohYes = "yeees"

// also fine
var ohYes2 string = "yay"

func main() {

	// the global variables defined in the module file are available
	// and can be shadowed inside a function in the same file / module
	fmt.Println(ohYes)

	// declaration +  assignment
	// best practice is to let go infer the type
	var favBook = "Lord of the rings"
	fmt.Println(favBook) // "Lord of the rings"
	// show the type of the variable
	fmt.Printf("%T\n", favBook) // string

	// individual characters can be accessed with array syntax
	secondLetterOrd := favBook[1] // 111
	// cast to string
	secondLetter := string(favBook[1])         // o
	fmt.Println(secondLetterOrd, secondLetter) // 111 o

	// reassignment
	favBook = "Rambo the novel"
	fmt.Println(favBook)
	// go is typed, favBook was initialized with a string, so
	// it can't be reassigned to a number
	// favBook = 12;

	// explicit type declaration and assignment
	var bigbigNumber complex64 = 3
	fmt.Println(bigbigNumber)
	var thirdBook string
	thirdBook = "Brothers Karamasov"
	fmt.Println(thirdBook)

	var myAge int
	var isItTrue bool
	var someBook string

	// default values
	fmt.Println(myAge)    // 0
	fmt.Println(isItTrue) // false
	fmt.Println(someBook) // ""

	// create multiple variables (compound creation)
	var favNumber, favCandy = 123, "pure honey"
	// block creation
	var (
		favNumber2 = 123
		favCandy2  = "pure honey"
	)
	fmt.Println(favNumber, favCandy, favNumber2, favCandy2)

	// declaration and assignment without using var
	// this can only be used inside a function
	favAnimal := "tiger"
	fmt.Println(favAnimal)

	pet1, pet2, pet3 := "cat", "dog", "rat"
	fmt.Println(pet1, pet2, pet3)

	// so variables can be swapped without a tmp
	pet1, pet2 = pet2, pet1

	// variables can not be reassigned using :=
	// pet3 := "other pet"

	// UNLESS multiple variables are assigned and at least one is new
	pet4, pet3 := "zebra", "still rat"
	fmt.Println(pet4, pet3)

	// constants (can't be reassigned)
	const myName = "pkro"

	// type casting
	var a int = 10
	var b int8 = 3
	var c = a + (int(b))

	fmt.Println(c)
}
