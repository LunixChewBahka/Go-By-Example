package main

import (
	"fmt"
	"github.com/lunixchewbahka/Go_By_Example/99_Scope/02_package_scope/math"
)

// Global var or package level
var iron int = 26

func main() {

	gold := 79

	fmt.Println(iron) // This works
	fmt.Println(gold) // This works

	xs := []float64{11, 22, 33, 44}
	avg := math.Average(xs)
	fmt.Println(avg)

	// fmt.Println(characters.Sheriff())
}

func anotherFunction() {

	fmt.Println(iron) // This will work because iron is decl. && init globally
	// fmt.Println(gold) // This doesn't work, gold won't reach not in scope
}
