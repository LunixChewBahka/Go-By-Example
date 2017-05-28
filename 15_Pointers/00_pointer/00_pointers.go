package main

import (
	"fmt"
)

/*

Go supports pointers, allowing you to pass references to values and records within your program.

*/

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

/*
When we call a function that takes an argument, that argument is copied
to the function.
Using func zero(params) will not modify the original x variable in the main function. If we wanted to change that we have to use pointers.
*/
func zero(x int) {
	x = 0
}

// function which has a var called xPtr which is a pointer to int
func zero_ptr(xPtr *int) {
	// acknowledges that the programmer passed in the address of x to the
	// &x carries the value "5" with it and then passes it to xPtr
	// function zero_ptr; function makes use of the address to point to the actual value of x in memory which is at the moment, "0"
	// "*" character (asterisk) is used to dereference pointer variables
	// dereferencing comes in then changes the value of x from "0" to "5"
	// "store the int 0 in the memory locatio xPtr refers to"
	*xPtr = 0
	// xPtr = 0
	// we will get an error because xPtr is not an "int", it's a "*int"; which
	// can only be given another *int.
}

func ptr_new(newPtr *int) {
	*newPtr = 1
}

func square(x *float64) {
	*x = *x * *x
	//*x = 2.5
}

func swap_num(a, b *int) {
	*a, *b = *b, *a
}

// func foo receives the address of ii (or &ii) then points to p which contains
// the value of *p + 3
func foo(p *int) {
	*p += 3 // *p = *p + 3 or *p = 9 + 3; which equates to 12
}

type Cat struct {
	Name  string
	Lives int
	Age   float32
}

func killCat(c *Cat) {
	c.Lives--
}

func foo2(a *[4]int) {
	a[0] = 3
	a[1] = 222
	a[2] = 333
	a[3] = 4444
}

func max(a *int, b *int, p **int) {
	if *a > *b {
		*p = a
	} else {
		*p = b
	}
}

func main() {

	/*

		Another way to get a pointer is to use the built-in 'new' function. 'new'
		takes a type as an argument, allocates enough memory to fit a value of that
		type and returns a pointer to it. We do not need to delete every instance of 'new' within a program as to some other programming languages wherein you need to take care of every 'new' instance by 'delete'ing them. Golang has a builtin garbage collection feature; which means memory is free'd / cleaned up automatically for the developer when it detects that nothing refers to it anymore.

	*/
	newPtr := new(int)

	x := 5
	i := 1

	problem_x := 1.5

	a := 10
	b := 20
	fmt.Printf("Value Before for a: %d and b: %d\n", a, b)
	swap_num(&a, &b)
	fmt.Printf("Values After for a: %d and b: %d\n", a, b)

	square(&problem_x) // 2.25?
	fmt.Println(problem_x)

	ptr_new(newPtr)
	fmt.Println("Using the 'new' builtin func:", *newPtr)
	fmt.Println("Address of 'i := 1' =>", &i)

	zero(x)
	fmt.Println("Not using a pointer:", x) // x is still the same :(
	fmt.Println("Address of 'x := 5' =>", &x)

	// refer to the location in memory of x instead of the actual value
	// "&" operator is used to locate the address of a variable.
	// &x return a *int (pointer to an int)
	// &x is main and xPtr in zero_pte refer to the same memory location
	zero_ptr(&x)
	fmt.Println("Using a pointer:", x) // x should now be 0
	fmt.Println("Address of 'x := 5' =>", &x)

	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	var xx int = 5
	var yy *int // pointer to int
	yy = &xx    // &x returns pointer
	fmt.Println("The value of the address the pointer is point to:", yy)

	zz := *yy // (dereference y to get the value of x)A
	fmt.Println("Dereferenced Value / Get the value on that Address:", zz)
	xx = 7
	fmt.Println("Modify the original Value", xx)
	zz = *yy
	fmt.Println("Updated value of xx:", zz)

	var ii = 9
	foo(&ii) // pass the address of ii not the value
	jj := ii // should be 12 right?
	fmt.Println("Value is passed to a function:", jj)

	cc := Cat{Name: "Arki", Lives: 9, Age: 4.3}
	killCat(&cc)
	//xc := cc.Lives
	fmt.Println(cc.Lives)

	var nums [4]int
	foo2(&nums)
	xn := nums[3] // 3?
	fmt.Println(xn)

	gg := 3
	hh := 5
	var kk *int
	max(&gg, &hh, &kk)
}
