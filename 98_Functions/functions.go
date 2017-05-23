package main

import (
	"fmt"
)

func main() {

	doSomething()
	add(2, 4, 10)
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
