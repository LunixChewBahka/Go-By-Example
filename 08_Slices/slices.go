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

	// slices are flexible dynamic lists
	// slices are dynamic; you can delete / append
	// arrays have a fixed length

	slice_s := []string{"Yoh", "test", "init"}
	slice_r := make([]string, 3)

	// There is nothing in here output : []
	fmt.Println(slice_s)
	// There is some space being allocated : [    ]
	fmt.Println(slice_r)

	// simply adds value to each element in the slice
	slice_r[0] = "hello"
	slice_r[1] = "hi"
	slice_r[2] = "what's up"

	fmt.Println(slice_r[1])

	// example of a slice
	// slice_letter := []string{"a", "b", "c", "d"}

	// declare a variable called "sliceS" and give it a slice of bytes as type
	var sliceS []byte
	sliceS = make([]byte, 5, 5)
	// simply prints the entirety of the slice
	// [0 0 0 0 0]
	fmt.Println(sliceS)
	fmt.Println(len(sliceS))
	// kind of the same buy wait there's more....
	fmt.Println(cap(sliceS))

	slice_of_b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(slice_of_b[1:4])
	fmt.Println(string(slice_of_b[1:4]))
	fmt.Println(slice_of_b[:2])
	fmt.Println(string(slice_of_b[:2]))
	fmt.Println(slice_of_b[2:])
	fmt.Println(string(slice_of_b[2:]))

	slice_of_x := [...]string{"Лайка", "Белка", "Стрелка"}
	fmt.Println(slice_of_x)

	slice_of_d := []byte{'r', 'o', 'a', 'd'}
	slice_of_e := slice_of_d[2:]
	fmt.Println(slice_of_e)
	fmt.Println(string(slice_of_e))
	slice_of_e[1] = 'm'
	fmt.Println(slice_of_e)
	fmt.Println(string(slice_of_e))
	fmt.Println(slice_of_d)
	fmt.Println(string(slice_of_d))

	s = s[2:4]
	fmt.Println(s)
	s = s[:cap(s)]
	fmt.Println(s)

	increase_slice()
}

func increase_slice() {
	var s []byte
	// or s := make([]byte, 5)
	s = make([]byte, 5, 5)
	// some funky stuff going on around here
	// let us increase the capacity of a slice
	slice_of_t := make([]byte, len(s), (cap(s)+1)*2)
	//for i := range s {
	//slice_of_t[i] = s[i]
	//}
	// instead of using for, we have a builtin function which is:
	// func copy(dst, src []T) int
	copy(slice_of_t, s)
	s = slice_of_t
	fmt.Println(s)

	// append using the long version
	slice_of_p := []byte{2, 3, 5}
	slice_of_p = AppendByte(slice_of_p, 7, 11, 13)
	fmt.Println(slice_of_p)

	// append using the builtin function
	// makes sense! ^________^
	// this one adds a slice to the end of an existing one
	slice_of_p2 := make([]int, 1)
	fmt.Println(slice_of_p2)
	slice_of_p2 = append(slice_of_p2, 1, 2, 3, 4)
	fmt.Println(slice_of_p2)

	// To append one slice to another, use "..." to expand the second argument to
	// list of args; in short slice + slice
	a_slice := []string{"John", "Paul"}
	b_slice := []string{"George", "Ringo", "Pete"}
	a_slice = append(a_slice, b_slice...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println(a_slice)
}

// long version of an already existing one from the builtin library
// func append([]type, ...type) []type
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
