package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000
	const z = 5

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Println(math.Sin(z))

	const goldenRatio float64 = 1.618034
	fmt.Println(goldenRatio)

	// Constant Enumeration
	const (
		First = iota
		Second
		Third
	)

	fmt.Println(First, Second, Third)
}
