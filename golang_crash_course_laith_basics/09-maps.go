package main

import "fmt"

func main() {
	// maps are unordered like in most other languages

	// create empty mape: make(map[key-type]val-type)
	cart := make(map[string]int)
	cart["milk"] = 3
	cart["cheese"] = 2
	fmt.Println(cart) // map[cheese:2 milk:3]

	// default values still work
	cart["someNewItem"] += 3
	fmt.Println(cart) // map[cheese:2 milk:3 someNewItem:3]

	// accessing values
	fmt.Println(cart["chees"])

	// check if a value exists and assign it
	// the second value is a boolean that is true if the item was found in the map
	milkInStore, foundItem := cart["milk"]
	if foundItem {
		fmt.Println(milkInStore)
	}

	// delete items
	delete(cart, "cheese")

}
