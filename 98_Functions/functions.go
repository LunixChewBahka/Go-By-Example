package main

import (
	"fmt"
)

func main() {

	doSomething()
	add(2, 4, 10)

	// result then receives the returned value from the func
	result := superComplexCalculation()
	fmt.Println(result)
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
func superComplexCalculation() int {
	// this function is executed then...
	e := 65*2/(2+5) - 6*2
	// e is returned as an int
	return e
}
