package main

import (
	"fmt"
)

// initialize firstName and lastName variables under Person struct
type Person struct {
	firstName string
	lastName  string
}

// An attempt to change the person's first name
func changeName(p *Person) {
	p.firstName = "Bob"
	p.lastName = "Ong"
}

func main() {
	// supply values to Person struct
	person := Person{
		firstName: "Alice",
		lastName:  "Dow",
	}

	fmt.Println("The original value:", person)
	// Check the location of person in memory
	fmt.Printf("Memory Address of struct Person: %p\n", &person)
	// pass the address in memory of var "person" which takes the value of "Person{} which is a struct"
	changeName(&person)

	fmt.Println("After the changeName:", person)
}
