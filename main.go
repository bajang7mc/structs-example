package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Hardy",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 10140,
		},
	}

	jim.updateFirstName("Jimmy")
	jim.print() //value is not update.
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v+\n", p)
}
