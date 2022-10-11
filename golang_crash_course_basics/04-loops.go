package main

import "fmt"

func main() {

	// only one loop construct (no while or do)

	// classic for loop
	for i := 0; i < 10; i++ { // paranthesis are always required, even for just one line
		fmt.Println(i)
	}

	// while
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}

	// foreach (also see range)
	var ids = []int{12, 33, 53, 6324, 4}

	for idx, id := range ids {
		fmt.Println(idx, id)
	}

	// "infinite" loop with break
	x := 0
	for {
		if x > 2 {
			break
		}
		x++
	}

}
