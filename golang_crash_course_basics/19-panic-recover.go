package main

import "fmt"

func main() {
	vals := []int{1, 2, 3}
	v := vals[10] // will cause "panic: runtime error: index out of range [10] with length 3"
	fmt.Println(v)
}
