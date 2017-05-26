package main

import (
	"fmt"
)

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	slice := array[:2]

	fmt.Println(slice, array) // [1 2] [1 2 3 4 5]

	// makes a copy of the array and add more memory as needed under the hood
	slice = append(slice, 8)

	fmt.Println(slice, array) // [1 2 8] [1 2 3 4 5]

	slice[0] = 0
	fmt.Println(slice, array) // [0 2 8] [1 2 3 4 5]

	/*

		Channels
		-> Buffered or unbuffered
		-> Buffering changes the sunchronisation model
		-> For signals use empty structure to save memory and show intent
			 c3 := make(chan struct{})

	*/
	c := make(chan int)
	c2 := make(chan int, 3)

}
