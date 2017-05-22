package main

import (
	"fmt"
)

func main() {

	// emp: [  ]
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// set: [a b c]
	fmt.Println("set: ", s)
	// get: c
	fmt.Println("get: ", s[2])

	// len: 3
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	// apd: [a b c d e f]
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	// simple copies all the contents from slice "s" to a new slice "c"
	copy(c, s)
	// prints the contents of the slice
	fmt.Println("cpy: ", c)

	// s[2:5] starts from index[2] which is the letter "c" up until index[5] but not including index[5]; therefore: the output should be [c d e]
	l1 := s[2:5]
	fmt.Println("sl1: ", l1)

	l2 := s[:5]
	// [a b c d e]
	fmt.Println("sl2: ", l2)

	l3 := s[2:]
	// [c d e f]
	fmt.Println("sl3: ", l3)

	t := []string{"g", "h", "i"}
	// [g h i]
	fmt.Println("dcl: ", t)

	// multi-dimensional arrays; length of the inner slices can vary
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	values := [][]int{}

	// These are the first two rows.
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	// Append each row to the two-dimensional slice.
	values = append(values, row1)
	values = append(values, row2)

	// Display first row.
	fmt.Println("Row 1")
	fmt.Println(values[0])

	// Display second row.
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Access an elements.
	fmt.Println("First element")
	fmt.Println(values[0][0])

	// Display entire slice.
	fmt.Println("Values")
	fmt.Println(values)
}
