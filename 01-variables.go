package main

import "fmt"

func main() {
	// declaration +  assignment
	// best practice is to let go infer the type
	var favBook = "Lord of the rings"
	fmt.Println(favBook)
	// reassignment
	favBook = "Rambo the novel"
	fmt.Println(favBook)
	// go is typed, favBook was initialized with a string, so
	// it can't be reassigned to a number
	// favBook = 12;

	//
	var otherBook string = "War and peace"
	fmt.Println(otherBook)
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
	favAnimal := "tiger"
	fmt.Println(favAnimal)

	pet1, pet2, pet3 := "cat", "dog", "rat"
	fmt.Println(pet1, pet2, pet3)

	// variables can not be reassigned using :=
	// pet3 := "other pet"

	// UNLESS multiple variables are assigned and at least one is new
	pet4, pet3 := "zebra", "still rat"
	fmt.Println(pet4, pet3)

	// constants
	const myName = "pkro"

}
