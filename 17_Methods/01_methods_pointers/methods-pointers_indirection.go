/*
Pointer and pointer indirection

Comparing the previous two programs, you might have noticed the functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v)	//Compile error!
ScaleFunc(&v)	// OK

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5) // This is OK
p := &v
p.Scale(10) // This is also OK

For the statement v.Scale(5) even though v is a vlue and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

*/
package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{3, 4}
	p.Scale(2)
	ScaleFunc(p, 10)

	fmt.Println(v, p)
}
