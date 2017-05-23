package main

import (
	"fmt"
)

func main() {

	doSomething()
	add(2, 4, 10)

	// result then receives the returned value from the func
	result_int, result_string, result_slice, sliceFloats := superComplexCalculation()
	fmt.Println(result_int, result_string, result_slice, sliceFloats)
}

func doSomething() {
	fmt.Println("Performing the doSomething() function")
}

/*
func function_name(var <type> ....params) <return type> {
	// doSomething here
}
*/
func add(a int, b int, c int) {
	fmt.Println(a + b*c)
}

// function return
func superComplexCalculation() (int, string, []int, []float64) {
	// this function is executed then...
	sliceTest := []int{1, 2, 3}
	sliceFloats := []float64{2.222, 3.33, 4.12313, 8585.321312}
	e := 65*2/(2+5) - 6*2
	f := "This is a string"
	// e is returned as an int
	return e, f, sliceTest, sliceFloats
}
