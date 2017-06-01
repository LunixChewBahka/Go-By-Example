package main

import (
	"fmt"
	"math"
)

/*

Go support methods defined on struct types

Methods and functions aswell
Remember: a method is just a function with a special receiver argument

can also declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
*/
type MyFloat float64

func (f MyFloat) Abs1() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// create a new type variable of Vertex
type Vertex struct {
	X, Y float64
}

/*

a function that has a speshul receiver arg a.k.a a "method"
The receiver appears in ti's own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.

can also be written this way to satisfy that a function == method:
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

*/
func (v Vertex) Abs() float64 {
	//square both values and add them up then return the sqrt of the total val
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

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

	// v takes the supplied values for X and Y which is part of our Vertex struct
	v := Vertex{3, 4}
	// pass the stuff that v took and then pass it to our method Abs()
	// expects to receive a float64 value after the operation
	fmt.Println(v.Abs())
	// can also be written this way to satisfy that a function == method:
	// fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs1())

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
