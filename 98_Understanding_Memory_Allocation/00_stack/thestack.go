package main

//import (
//"fmt"
//)

/*

The Stack

^ Represents the "state" of execution
^ Composed of stack frames
^ Each frame holds a function call and lcoal variables
^ Last in First Out

What's wrong with it?
- One stack = one execution state
- Stacks usually have a pre-defined capacity
- need to duplicate the entire model for concurrency (bad for concurrency)
---- if you call too many functions and go deep you are going to have a stackoverflow

Go's Solution:
+ One heap, one stack per goroutine
+ Elastic stacks
+ Stacks borrow memory from the heap

*/

// very first frame in the stack
func main() {
	a, b := 1, 2
	// this is only for demonstration purposes
	result := level1(a, b)
	println(result)
}

// 2nd frame on the stack
func level1(c, d int) int {
	// add another argument to satisfy level2
	e := 3
	return level2(c, d, e)
}

// last frame that was added on the stack
func level2(f, g, h int) int {
	// 1 + 2 + 3; returns "6"
	return f + g + h
}
