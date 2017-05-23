package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		// should execute 3x since i := 1; if i := 0 then 4x
		fmt.Println(i)
		i = i + 1
		// or i += 1
	}

	// this chunk will execute 3x
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		// will print once
		fmt.Println("loop")
		// after printing, it will break out of the for loop
		break
	}

	// will loop 6x
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		// but will only print the ones that are not even
		// in this case 1, 3, 5
		fmt.Println(n)
	}

	// for loop the traditional way
	fmt.Println("\nTraditional Way: ")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nPreferred way of typing for in golang: ")
	myName := "Julian"
	/*
		for <identifier>, <variable> := range <iterable> {
			// do something about it
		}
	*/
	for _, char := range myName {
		// char will only return the ASCII code; a workaround is to convert it
		// to type string
		// fmt.Println(char, " ")
		fmt.Println(string(char) + " ")
	}
	fmt.Println("")

	/* Be careful of handling for loops
	Example: (infinite for loop)
	for {
		// do something about it
	}
	*/

	// controlled for loop
	z := 0

	// should loop 6x
	for z < 5 {
		fmt.Println("Go Language")
		z++
		// but there was a nested condition declared and initialized inside the loop which halts / skips the iteration once i is exactly equal to 3.
		if z == 3 {
			// break out of the for loop once the condition is met.
			break
		}
	}
}
