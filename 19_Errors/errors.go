package main

import (
	"errors"
	"fmt"
)

/*

In go it's idiomatic to communicate errors via an explicit, serparate return
value. This contrasts with the exceptions used in languagues lije Java and Ruby and the overloaded single result / error valye sometimes used in C. Go' approach makes it easy to tsee which functions return errors and to handle them using the same langauge constructs employed for any other, non-error tasks.

*/

// By convention, errors are the last return value and have type error,
// a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {

		// errors.New constructs a basic error value with the given error msg.
		return -1, errors.New("can't work with 42")

	}

	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

/*

It's possible to use custom types as errors by implementing the Error() method on them. Here's a vriant on the example above that uses a custom type to explicitly
represent an argument error.

*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "cant' work with it"}

	}

	return arg + 3, nil

}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
