package main

import (
	"fmt"
)

var x int

func f1() {
	fmt.Println(x)
}

func f2() {
	fmt.Println(x)
}

func main() {
	fmt.Println(y)            // invoked from f1.go
	fmt.Println(global_x - 3) // var invoked from f1.go
	f1()
	f2()
	fmt.Println("Hello World")
	{
		i := 5
		fmt.Println(i)
	}
}
