package main

import (
	"fmt"
)

func main() {
	var a [5]int
	// emp: [0 0 0 0 0]
	fmt.Println("emp:", a)

	a[4] = 100
	// set: [0 0 0 0 100]
	fmt.Println("set:", a)
	// get: 100
	fmt.Println("get:", a[4])

	// len: 5
	fmt.Println("len:", len(a))

	// initialize and array which has 5 elements {1 .. 5}
	b := [5]int{1, 2, 3, 4, 5}
	// simple prints the entire array in one line
	fmt.Println("dcl:", b)

	/* declare and initialize a 2-D array
	an array of 2 arrays of 3 arrays of type int (mutli-dimensional array)
	*/

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
