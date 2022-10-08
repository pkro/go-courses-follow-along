package main

import "fmt"

func main() {
	type Animal struct {
		class  string
		age    int
		gender bool
	}

	var teddy = Animal{
		class:  "bear", // comma!
		age:    24,
		gender: true,
	}

	fmt.Println(teddy)
	fmt.Println(teddy.age)
	// structs are mutable
	teddy.age += 1
	fmt.Println(teddy.age)

	// skip field names; fields must be in order for this to work
	var leo = Animal{"lion", 12, false}
	fmt.Println(leo)

	// Not assigning property values
	var lalo = Animal{}
	fmt.Println(lalo) // { 0 false}

	// anonymous struct
	var tuco = struct {
		class  string
		age    int
		gender bool
	}{
		class:  "Penguin",
		age:    1,
		gender: false,
	}

	fmt.Println(tuco)

	// Pointers are fine with structs
	pTuco := &tuco
	//pointers are automatically dereferenced.
	fmt.Println(pTuco.age)

}
