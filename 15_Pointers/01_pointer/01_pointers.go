package main

import (
	"fmt"
)

//func updateValue(val int) {
//val = val + 100
//}

/*

change int to *int ; dereferencing

A pointer points to a memory address lcoation and allows us to change the value.
we can call the value anything we want inside a function; it doesn't matter

*/
func updateValue(someVal *int) {
	*someVal = *someVal + 100
}

// Another way is by using the builtin keyword "new" instead of "&"
func updateValue_2(whateverVal *int, whateverVal2 *float64) {
	*whateverVal = *whateverVal + 100
	*whateverVal2 = *whateverVal2 + 1.75
}

// pointer using structs
type Stock struct {
	high  float64
	low   float64
	close float64
}

// modify the items in stock struct
func modifyStock(stock *Stock) {
	stock.high = 475.10
	stock.low = 400.15
	stock.close = 450.75
}

func main() {

	// example 1
	val := 1000

	// example 2
	val_0 := 2000
	val_1 := new(float64)

	// example 3
	goog := Stock{454.43, 421.01, 435.29}
	// this doesn't change the value :p
	// I was also expecting a "1100" instead of "1000"
	//updateValue(val)
	//fmt.Println("val:", val)

	/*
		To fix this we need to pass the val variable by reference
	*/
	updateValue(&val)
	fmt.Println("Example #1: ")
	fmt.Println("val:", val)
	// output now should be : "1000"

	updateValue_2(&val_0, val_1)
	fmt.Println("\nExample #2: ")
	fmt.Println("val_0:", val_0)
	fmt.Println("val_1:", *val_1)

	fmt.Println("\nExample #3: ")
	fmt.Println("Original Stock Data:", goog)
	modifyStock(&goog)
	fmt.Println("Modified Stock Data:", goog)
}
