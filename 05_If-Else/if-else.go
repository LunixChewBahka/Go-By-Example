package main

import "fmt"

func main() {

	// this block simple checks if 7 is and odd or even
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// checks if 8 is divisible by 4
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// evaluates the number 9;
	if num := 9; num < 0 {
		// 9 is definitely not a negative #
		fmt.Println(num, "is negative")
	} else if num < 10 {
		// 9 is less than 10 which is true; caught condition
		fmt.Println(num, "has 1 digit")
	} else {
		// 9 doesn't have any multiple digits
		fmt.Println(num, "has multiple digits")
	}

	condition := 5 > 2

	if condition {
		fmt.Println("5 is inded greater than 2")
	} else {
		fmt.Println("This condition is false!")
	}

	// && AND example; statement && statement
	/*
		t	f	r
		0	0   0
		0	1	0
		1	0	0
		1	1	1
	*/
	if 6 > 3 && 5 > 8 {
		fmt.Println("Workded")
	} else {
		fmt.Println("Didn't work")
	}

	// || OR statement || statement
	/*
		t	f	r
		0	0	0
		0	1	1
		1	0	1
		1	1	1
	*/
	if 6 > 3 || 5 > 8 {
		fmt.Println("Worked")
	} else {
		fmt.Println("Didn't work")
	}
}
