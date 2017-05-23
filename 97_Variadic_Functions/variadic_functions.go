package main

import (
	"fmt"
)

func main() {

	r := add(1, 22, 333, 4444, 5555, 666666, 7777777)
	fmt.Println(r)

}

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
