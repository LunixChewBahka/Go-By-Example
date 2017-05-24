package mzzain

import (
	"fmt"
)

func dosomething() {
	fmt.Println("performing the dosomething() function")
}

/*
func function_name(var <type> ....params) <return type> {
	// dosomething here
}
*/
func add(a int, b int, c int) {
	fmt.Println(a + b*c)
}

// function return
func supercomplexcalculation() (int, string, []int, []float64) {
	// this function is executed then...
	slicetest := []int{1, 2, 3}
	slicefloats := []float64{2.222, 3.33, 4.12313, 8585.321312}
	e := 65*2/(2+5) - 6*2
	f := "this is a string"
	// e is returned as an int
	return e, f, slicetest, slicefloats
}

func plus(a int, b int) int {

	return a + b

}

func plusplus(a, b, c int) int {

	return a + b + c

}

func main() {

	dosomething()
	add(2, 4, 10)

	// result then receives the returned value from the func
	result_int, result_string, result_slice, slicefloats := supercomplexcalculation()
	fmt.Println(result_int, result_string, result_slice, slicefloats)

	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

}
