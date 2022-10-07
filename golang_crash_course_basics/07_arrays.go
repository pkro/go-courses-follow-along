package main

import (
	"fmt"
	"reflect"
)

func main() {
	// initialize without values
	var stuff [3]string
	stuff[1] = "hey"
	fmt.Println(stuff) // ["" "hey" ""]

	// initialize & assign array
	purchases := [5]float32{9.99, 2, 4, 95, 4}
	fmt.Println(purchases)

	// size can be skipped if values are assigned during initialization
	moreStuff := []int{1, 2, 3}
	fmt.Println(moreStuff)

	// error - array length is fixed so no new entries can be added:
	//purchases[5] = "hey"

	// arrays are not read only (if that needs to be said)
	purchases[len(purchases)-1] += 10
	fmt.Println(purchases) // {9.99, 2, 4, 95, 4}

	// get a range of items like in python; also see 08-slices.go
	fmt.Println(purchases[1:3])                 // [2 4]
	fmt.Println(reflect.TypeOf(purchases[1:3])) // []float32

	// mixed value type array using interface
	var mixed [3]interface{}
	mixed[0] = 1
	mixed[1] = false
	mixed[2] = "a string!"
	fmt.Println(mixed)

	// still only one loop type in go
	for i := 0; i < len(purchases); i++ {
		fmt.Println(purchases[i])
	}

	// foreach-like
	for index, itr := range purchases {
		fmt.Print(index, " : ", itr, "\n")
	}

	// if we use only one variable in the for loop,
	// it by default refers to the value in the container.
	for val := range purchases {
		fmt.Print(val, " ")
	}
}
