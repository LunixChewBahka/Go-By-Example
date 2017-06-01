package main

/*
Interfaces in Go

Interfaces make the code more flexible, scalable and it's a way to achieve
polymorphism in Golang. Instead of requiring a particular type, interfaces allow to specify that only some behavious is needed. Behaviour is defined by set of
methods

*/

import (
	"fmt"
)

type X int64

/*
No particular implementation is enforced. It's enough that type defines methods with desired names and signatures (input and output parameters) to say it implements (satisfies) an interface:
*/
type I interface {
	//f1(name string)
	//f2(name string) (error, float32)
	//f3() int64
	M()
}

type I1 interface {
	M1()
}

type I2 interface {
	M2()
}

type T struct {
	name string
}

type T1 struct{}
type T2 struct{}

func (T) M1() { fmt.Println("T.M1") }
func (T) M2() { fmt.Println("T.M2") }

func (T1) M() { fmt.Println("T1.M") }
func (T2) M() { fmt.Println("T2.M") }

func f(i I) { i.M() }

func f1(i I1) { i.M1() }
func f2(i I2) { i.M2() }

// Type T satisfies interface I defined above. Values of type T can be f.ex. passed to any function accepting I as a parameter
func (X) f1(name string) {
	fmt.Println(name)
}

func (X) f2(name string) (error, float32) {
	return nil, 10.2
}

func (X) f3() int64 {
	return 10
}

func (t T) M() string {
	return t.name
}

func Hello(i I) {
	fmt.Printf("Hi, my name is %s\n", i.M())
}

func main() {
	Hello(T{name: "Jar"})
	// should print out "Hi, my name is Jar"
	t := T{}
	f1(t)
	f2(t)

	f(T1{})
	f(T2{})
}
