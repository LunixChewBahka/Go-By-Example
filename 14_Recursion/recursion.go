package main

import (
	"fmt"
	"log"
	"time"
)

/*

Go support recursive functions. Here's a classic factorial example.

*/

// This fact functions calls itself until it reaches the base case of fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prog_time() {

	prog_start := time.Now()

	prog_elapse := time.Since(prog_start)
	log.Printf("Everyting took %s\n", prog_elapse)

}

func main() {

	prog_time()

	fib_nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, z := range fib_nums {
		fmt.Println(fact(z))
	}

}
