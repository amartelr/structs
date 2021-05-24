package main

import "fmt"

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

	me := person{
		firstName: "Alfredo",
		lastName:  "Martel",
		contact: contactInfo{
			email:   "alfredo.martel@gmail.com",
			zipCode: 0,
		},
	}

	me.updateNameByValue("Alf")
	me.print()

	me.updateNameByRef("Alf")
	me.print()

}

// Don't new Pointers: slices, maps, channels, pointers, functions
func (p person) updateNameByValue(newFirstName string) {
	p.firstName = newFirstName
}

// Use Pointers to change things in functions for: int, float, string, bool, structs
func (p *person) updateNameByRef(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
