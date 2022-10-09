package main

import "fmt"

type Animal struct {
	class  string
	age    int
	gender bool
}

// it is idiomatic to create struct instances with factory methods prefixed by New
func NewAnimal(class string, age int, gender bool) (Animal, error) {
	if age >= 0 && class != "" { // example data validation
		return Animal{
			class:  class,
			age:    age,
			gender: gender,
		}, nil
	}
	return Animal{}, fmt.Errorf("class empty or age < 0")
}

// for large data structures we can of course return also a pointer;
// the compiler will allocate the memory accordingly so there's no
// computational overhead during execution
// the usage is the same, e.g. animal := NewAnimalP()
func NewAnimalP() (*Animal, error) {
	// return empty animal for brevity
	return &Animal{}, nil
}

func main() {
	// normal instantiation
	var teddy = Animal{
		class:  "bear", // comma!
		age:    24,
		gender: true,
	}

	fmt.Println(teddy)
	fmt.Println(teddy.age)

	// instantiation using factory defined above
	var mauki, err = NewAnimal("cat", 3, true)
	if err == nil {
		fmt.Println(mauki)
	}

	// structs are mutable
	teddy.age += 1
	fmt.Println(teddy.age)

	// skip field names; fields must be in order for this to work
	var leo = Animal{"lion", 12, false}
	fmt.Println(leo)

	// Not assigning property values
	var lalo = Animal{}
	fmt.Println(lalo) // { 0 false}

	// anonymous struct
	var tuco = struct {
		class  string
		age    int
		gender bool
	}{
		class:  "Penguin",
		age:    1,
		gender: false,
	}

	fmt.Println(tuco)

	// Pointers are fine with structs
	pTuco := &tuco
	//pointers are automatically dereferenced.
	fmt.Println(pTuco.age)

}
