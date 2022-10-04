package main

import "fmt"

func main() {
	purchases := [5]float32{9.99, 2, 4, 95, 4}

	// get a slice of an array, like to python list style
	mySlice := purchases[:] // all items
	fmt.Println(mySlice)

	// syntax is [startIdx:endIdx+1] like in python
	fmt.Println(purchases[1:3]) // [2 4]

	// omitting means from the start / to the end
	fmt.Println(mySlice[:2]) // [9.99 2]

	mySlice = append(mySlice, 99)
	fmt.Println(mySlice) // [9.99 2 4 95 4 99]

	// inline creation of slice
	anotherSlice := []float32{1, 2, 3, 4, 5}
	fmt.Println(anotherSlice)

	// combining slices; like the JS spread operator, just on the right side
	combined := append(mySlice, anotherSlice...)
	fmt.Println(combined)
}
