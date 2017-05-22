package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2 // 2 which is an int
	// expected output should be:
	// Write 2 as two.
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// simple checks what day it is of the week (now)
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	command := "close"

	switch command {
	case "create":
		fmt.Println("Creating...")
	case "open":
		fmt.Println("Opening...")
	case "close":
		fmt.Println("Closing...")
	default:
		fmt.Println("Unrecognized command")
	}
	// result should print "Closing..."
}
