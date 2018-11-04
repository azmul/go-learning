package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// Declaration of Struct
	person1 := person{lastName: "Hossain", firstName: "Azmul"}
	fmt.Println(person1)

	// Update Struct 1
	var person2 person
	person2.firstName = "Azmul"
	person2.lastName = "Hossain"
	// fmt.Printf("%+v", person2)

	// Update Struct 2
	person3 := person{
		firstName: "Azmul",
		lastName:  "Hossain",
		contact: contactInfo{
			email:   "azmulg@gmail.com",
			zipCode: 1200,
		},
	}
	person3.updateName("Nazmul")
	person3.print() // call print function
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() { // receiver person
	fmt.Printf("%+v", p)
}
