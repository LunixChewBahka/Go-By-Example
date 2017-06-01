package main

import (
	"fmt"
)

type I interface {
	M()
	method1()
}

type T struct{}

type T1 struct{}

//type T2 struct{}

func (T) method1() {}

func (T1) M() { fmt.Println("T1.M") }

//func (T2) M() { fmt.Println("T2.M") }

func f(i I) { i.M() }

func main() {
	f(T1{})
	//f(T2{})

	var i I = T{}
	fmt.Println(i)

}
