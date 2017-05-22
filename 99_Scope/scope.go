package main

import (
	"fmt"
)

// Global var or package level
var iron int = 26

func main() {

	gold := 79

	fmt.Println(iron) // This works
	fmt.Println(gold) // This works

}

func anotherFunction() {

	fmt.Println(iron) // This will work because iron is decl. && init globally
	// fmt.Println(gold) // This doesn't work, gold won't reach not in scope
}
