package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	myAdder := adder()
	for i := 1; i <= 3; i++ {
		myAdder(i)
	}

	fmt.Println(myAdder(0)) // 6
}
