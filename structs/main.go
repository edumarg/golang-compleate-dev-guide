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
	contact   contactInfo
}

func main() {
	edu := person{
		lastName:  "Marg",
		firstName: "Edu",
		age:       40,
		contact: contactInfo{
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
	ber.contact = contactInfo{phone: "555-4444-4444", email: "ber@mail.com"}
	fmt.Printf("%+v", ber)
}
