package main

import (
	"fmt"
)

func main() { // emp: [  ]
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
	row2 := []int{5, 6, 7}

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

	// Create a two-dimensional array
	letters := [2][2]string{}

	// Assign all elements in 2 by 2 array
	letters[0][0] = "a"
	letters[0][1] = "b"
	letters[1][0] = "c"
	letters[1][1] = "d"

	// Display result
	fmt.Println(letters)

	// Nested string slices using range
	// Create an empty slice of slices
	//animals := [][]string{}

	//row1 := []string{"fish", "shark", "eel"}
	//row2 := []string{"bird"}
	//row3 := []string{"lizard", "salamander"}

	//animals = append(animals, row1)
	//animals = append(animals, row2)
	//animals = append(animals, row3)

	//// Loop over slices in animals
	//for i := range animals {
	//fmt.Printf("Row: %v\n", i)
	//fmt.Println(animals[i])
	//}
	phrase1 := "The quick brown fox"
	phrase2 := " jumps over the lazy dog"

	fmt.Println(phrase1 + phrase2)

	// Slicing strings
	// 3 is non-inclusive
	aSliceOfText := phrase1[1:3]
	fmt.Println(aSliceOfText)

	noun1 := phrase1[16:]
	fmt.Println(noun1)

	noun2 := phrase2[21:]
	fmt.Println(noun2)

	character := phrase1[0]
	// it returns the ascii code 84 instead of "T"
	fmt.Println(character)
	// convert the result to a string to get the actual "T"
	fmt.Println(string(character))
}
