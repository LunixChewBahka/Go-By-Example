package main

import (
	"fmt"
	"math"
)

type Circle struct {
	// instead of doing this,
	//x float64
	//y float64
	//z float64

	// do this,
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Embedded types - is a matter of "is a"; "has a" relationship
type Person struct {
	Name string
}

type Android struct {
	// using the other struct that we just initalized
	// Person person
	// but you can use another way without the type "Person"
	// this means Anroid is a Person rather than Android has a Person
	Person
	Model string
}

/*

You may have noticed that we were able to name the Rectangle's area method the
same thing as the Circle's area method. This was no accident. In both real life and in programming ,relationships like these are commonplace. Go has a wayof
making these accidental similarities explicit through a type known as an interface. Here is an example of a Shape interface

*/
type Shape interface {
	/*

		instead of defining fields, we define a "method set". A method set is a list of methods that a type must have in order to "implement" the interface.

	*/
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

/*

In our case both Rectangle and Circle have area methods which return float64 so both types implete the Shape interface. By itself this wouldn't be particularly useful, but we can use interface types as arguments to functions:

*/

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea1(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	//fmt.Println(l)
	w := distance(x1, y1, x2, y1)
	//fmt.Println(w)
	return l * w
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Area of a Circle, A=πr^2
// Version 1
func circleArea1(x, y, r float64) float64 {
	r_squared := r * r
	return math.Pi * r_squared
}

// Version 2 integrating structs
func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// Version 3 using structs + methods
// func <receiver> <function name()> <type>
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

//// Circumference of a Circle, C=2πr
//func circleCircumference(x, y, r float64) float64 {
//return 2 * math.Pi * r
//}

//// Diameter of a Circle, d=2r
//func circleDiameter(x, y, r float64) float64 {
//return 2 * r
//}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	// we can use this: var c Circle, or this one:
	// c := new(Circle)
	// we can also use this with supplied arguments:
	c := Circle{x: 0, y: 0, r: 5}
	//fmt.Println(c.x, c.y, c.r)
	// we could also leave off fields names if we know the order they were
	// defined
	// c := Circle{0, 0, 5}
	// 0, 0, 5
	// 0, 0, 10, 10

	r := Rectangle{0, 0, 10, 10}

	a := new(Android)
	b := new(Person)
	a.Talk()
	b.Talk()

	fmt.Println("Using Rectangle Ver #1:", rectangleArea1(rx1, ry1, rx2, ry2))
	fmt.Println("Using Rectangle Ver #2:", r.area())
	//fmt.Println(circleCircumference(cx, cy, cr))
	//fmt.Println(circleDiameter(cx, cy, cr))
	fmt.Println(math.Pi)

	fmt.Println("Using Circle Version #1:", circleArea1(cx, cy, cr))
	fmt.Println("Using Circle Version #2:", circleArea2(&c))
	//c.x = 10
	//c.y = 5
	//c.r = 0
	fmt.Println("Using Circle Version #3:", c.area())
}
