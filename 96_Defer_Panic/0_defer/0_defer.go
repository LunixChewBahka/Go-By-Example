package main

import (
	"fmt"
)

func main() {

	// this send the doSomething to the very end of the call stack
	// will be executed last within the package
	// leave this function for last
	defer doSomething()
	// call this
	doSomethingElse()

}

func doSomething() {
	fmt.Println("doSomething()")
}

func doSomethingElse() {
	fmt.Println("doSomethingElse")
}
