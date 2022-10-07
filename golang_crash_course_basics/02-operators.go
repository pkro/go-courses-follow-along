package main

import "fmt"

func main() {

	// arithmetic operators
	var num1 = 4
	var num2 = 3

	// + - * % as usual
	var sum = num1 / num2
	fmt.Println(sum)         // 1, as it's an integer
	fmt.Println(num1 / num2) // 1, as it's an integer

	var float1 float32 = 4
	var float2 float32 = 3
	var sumFloat = float1 / float2
	fmt.Println(sumFloat) // 1.3333334

	// fmt.Println(float1 / num2) // invalid operation: float1 / num2 (mismatched types float32 and int)

	// relational operators
	// < > == && || ! as usual

	// logical operators

	var a int = 10
	var b int = 3

	// are applied between each bit of the values, e.g.
	// a & b:
	// 10 = 1010
	// 3  = 0011
	// ---------
	// 2  = 0010

	fmt.Println(a & b)  // 2 - and
	fmt.Println(a | b)  // 11 - or
	fmt.Println(a ^ b)  // 9 - xor
	fmt.Println(a &^ b) // 8 - and not - only the first one has the bit set, but not the other
	fmt.Println(b &^ a) // 1 - so a &^ b has a different result than b &^ a, unlike the previous logical operators

	// bit shifting operators
	a = 8
	fmt.Println(a << 3) // 64 - 00001000 becomes 01000000
	fmt.Println(a >> 3) // 1  - 00001000 becomes 00000001

}
