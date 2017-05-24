package main

import (
	"fmt"
)

/*

Go support methods defined on struct types

*/
type rect struct {
	width, height int
}

// This are method has a reciver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here's an example of a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{width: 10, height: 5}

	// makes use of the code in line 12
	// example of how to call the 2 methods defined for our struct
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	/*

		Go automatically handles conversion between values and pointer for method
		calls. You may want to use a pointer receiver type to avoid copying on
		method or to allow the method to mutate the receiving struct.

	*/
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

}
