package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	age       uint8
	// contact   contactInfo
	contactInfo // No need field name,  you can use only the name of the type.
}

func main() {
	edu := person{
		lastName:  "Marg",
		firstName: "Edu",
		age:       40,
		contactInfo: contactInfo{
			email: "edu@mail.com",
			phone: "555-5555-5555",
		},
	}
	fmt.Printf("%+v", edu)
	edu.lastName = "Smith"
	fmt.Printf("%+v", edu)
	var ber person
	fmt.Printf("%+v", ber)
	ber.age, ber.firstName, ber.lastName = 38, "Ber", "Spy"
	ber.contactInfo = contactInfo{phone: "555-4444-4444", email: "ber@mail.com"}
	fmt.Printf("%+v", ber)
	mike := person{firstName: "Mike", lastName: "Piazza", age: 50, contactInfo: contactInfo{email: "mike@mail.com", phone: "566-6666-6666"}}
	mike.print()
	// mikePointer := &mike // & is operator to give the memory address of the value variable is pointing to
	// mikePointer.updateName("Mikey")
	mike.updateName("Mikey") // shortcut to use pointers with a receiver function that uses as type  a pointer
	mike.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) { // * is an operator to give the value the memory address is point at
	(*pointerToPerson).firstName = newFirstName
}
