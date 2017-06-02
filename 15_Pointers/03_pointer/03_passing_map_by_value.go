package main

import (
	"fmt"
	"unsafe"
)

type MyType struct {
	Value1 int
	Value2 string
}

func main() {
	myMap := make(map[string]string)
	myMap["Bill"] = "Deoxys"
	pointer := unsafe.Pointer(&myMap)

	fmt.Printf("[1] Addr: %v Value: %s\n", pointer, myMap["Bill"])

	ChangeMyMap(myMap)
	fmt.Printf("[2] Addr: %v Value: %s\n", pointer, myMap["Bill"])

	ChangeMyMapAddr(&myMap)
	fmt.Printf("[3] Addr: %v Value: %s\n", pointer, myMap["Bill"])
}

func ChangeMyMap(myMap map[string]string) {
	myMap["Bill"] = "Joan of Arc"

	pointer := unsafe.Pointer(&myMap)

	fmt.Printf("[01_func]Addr: %v Value: %s\n", pointer, myMap["Bill"])
}

// Don't Do this, Just For Use In This Article
func ChangeMyMapAddr(myMapPointer *map[string]string) {
	(*myMapPointer)["Bill"] = "Jenny"

	pointer := unsafe.Pointer(myMapPointer)

	fmt.Printf("[02_func]Addr: %v Value: %s\n", pointer, (*myMapPointer)["Bill"])
}
