package main

import (
	"fmt"
)

/*

A goroutine is alightweight thread of execution

*/

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	/*
		Suppose we have a function call f(s). Here's how we'd call that in the
		usual way, running it synchronously
	*/
	f("direct")

	/*
		To invoke this function in a goroutine, use go f(s). This new goroutine will
		execute concurrently with the calling one.
	*/
	go f("goroutine")

	/*
		You can also start a goroutine for an anonymous function call.
	*/
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	/*
		Our two function calls are runing asnchronously in separate goroutines now,
		so execution falls through o here. This Scanln code requried we press a key
		before the program exits.
	*/
	var input string
	// will prompt to press <return>
	fmt.Scanln(&input)
	fmt.Println("done")
}
