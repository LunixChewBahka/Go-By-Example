package main

import (
	"fmt"
)

/*

Structs are types collections of fields. They're useful for grouping data
together to form records.

*/

type person struct {
	name string
	age  int
}

func main() {

	// can group different items with different types
	fmt.Println(person{"Bob", 20})

	// you can even label them
	fmt.Println(person{name: "Alice", age: 30})

	// supply name: "Fred" then age will be automatically initalized to "0"
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	// access fields using dot notation
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	// refers to "s" var which has the newly assigned person params / values
	// pointers are dereferenced automatically
	sp := &s
	fmt.Println(sp.age)

	// should change to 51
	// structs are mutable; meaning you can override or change them
	sp.age = 51
	fmt.Println(sp.age)

}
