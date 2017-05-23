package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// read the whole file at once
	b, err := ioutil.ReadFile("../names.txt")

	if err != nil {
		panic("The program failed at the error checking procedure")
	}

	for _, e := range b {
		fmt.Print(string(e))
	}
}
