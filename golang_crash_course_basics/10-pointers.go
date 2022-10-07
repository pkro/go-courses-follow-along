package main

import (
	"fmt"
)

func main() {
	// as in JS, when assigning a primitive variable to another
	// the value is copied
	var a = 1
	var b = a
	b += 1
	fmt.Println(a, b) // 1 2

	// the same is true for structs
	s := struct {
		someValue int
	}{
		someValue: 3,
	}

	var s2 = s
	s2.someValue += 10
	fmt.Println(s)  // {3}
	fmt.Println(s2) // {13}

	// array are assigned by reference / pointer,
	// so changing ar2 also changes ar as they are the
	// same object in memory
	ar := []int{1, 2, 3}
	ar2 := ar
	ar2[0] = 99
	fmt.Println(ar, ar2) // [99 2 3] [99 2 3]

	// if we want the same to happen with primitives (or structs), we can use pointers:
	var c = 1
	var d *int // d is a pointer to an int
	// assign d the memory location of c;
	// as we defined the type of the pointer above, go knows the end of the memory location
	d = &c
	fmt.Println(c, d) // 1 0xc000020118

	// dereference (access) / change the value at the pointers memory location
	*d++
	fmt.Println(c) // 2

	// passing references (memory address) to functions
	x := 5
	squareAdd(&x)
	fmt.Println(x) // 25
}

func squareAdd(p *int) {
	*p *= *p
}
