package main

import "fmt"

type Person struct {
	name, title string // shortcut for multiple properties of the same type
	age         int
}

// Adding a method to the struct;
// the struct members are accessible by p (like "this" in JS, name doesn't matter)

// value receiver (doesn't change anything)
func (p Person) greet() string {
	return fmt.Sprintf("Hey I'm %s and I'm %d years old", p.name, p.age)
}

// pointer receiver (can change values of the struct)
func (p *Person) celebrateBirthday() {
	p.age++
}

func main() {
	p := Person{
		name: "Joe",
		age:  30,
	}

	fmt.Println(p.greet()) // Hey I'm Joe and I'm 30 years old
}
