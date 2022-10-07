package main

import "fmt"

func main() {
	var ids = []int{12, 33, 53, 6324, 4}

	// foreach-like looping over array or map
	for idx, id := range ids {
		fmt.Println(idx, id)
	}

	// if the value is required, both key and value are necessary,
	// convention is to use "_" if one of them is unused
	for _, id := range ids {
		fmt.Println(id)
	}

	// value can be omitted if only key is required
	emails := map[string]string{"Bob": "bob@gmail.com", "Joe": "joe@gmail.com"}
	for name := range emails {
		fmt.Printf("Name is %s\n", name)
	}

}
