package main

import (
	"fmt"
)

/*

Go support anonymous functions, which can form closures. Anonymous functions are
useful when you want to define a function inline without having to nane it

*/

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
