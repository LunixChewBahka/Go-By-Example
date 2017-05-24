package main

import (
	"fmt"
)

/*
Variadic Functions

Variadic functions can be called with any number of trailing arguments. For
example, fmt.Println is a commin variadic function.

*/

// expects multiple arguments to be supplied to the parameters
// in this case there is only 1 parameter and that is the ...int
func add(nums ...int) int {
	// simply declares the variable total as type int
	var total int

	// the for loop golang uses
	// in range of the slice; "nums"
	for _, n := range nums {
		// adds all the integers that is part of the list
		// until we get the summation
		total += n
	}
	// return the total amount
	return total
}

func sum(nums ...int) {
	fmt.Print("values:", nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("=", total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	r := add(1, 22, 333, 4444, 5555, 666666, 7777777)
	fmt.Println(r)

}
