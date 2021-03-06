package main

import (
	"fmt"
)

/*

In Go, you define a method receiver to specify which struct to attach a certain
function in order to make it invoke-able as a method. For instance, func (d Dog)
is part which defines the method receiver in the following program:
*/

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("Woof Woof!")
}

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

// go:inline
func increment(inc int) {

	// Increment the "value of" inc
	inc++
	fmt.Println("inc:\tValue of [", inc, "]\tAddr Of [", &inc, "]")

}

func main() {

	d := &Dog{}
	d.Say()

	m := &Mutatable{0, 0}
	// output should be ${0 0}
	fmt.Println(m)

	m.StayTheSame()
	// as expected should stay the same
	// passing the value to a non-pointer receiver on line 25
	fmt.Println(m)

	m.Mutate()
	// output should the the values: &{5 7}
	// passing the value to a pointer @ line 30
	fmt.Println(m)

	// Declare variable of type int with a value 10
	count := 10

	// Display the "value of" and "address of " count
	fmt.Println("count:\tValue Of [", count, "]\tAddr Of [", &count, "]")

	// Pass the "value of" the count
	increment(count)
	fmt.Println("count:\tValue Of [", count, "]\tAddr Of [", &count, "]")
}
