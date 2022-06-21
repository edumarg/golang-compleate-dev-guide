package main

import "fmt"

// Create a new type of "deck"
// which is a slice of strings

type deck []string

// receiver function: a function that receives a variable.
// func (<variable name> <variable type>) <function name> {logic}

func (d deck) print() { // receiver function that receives any variable type deck and use it inside the function logic
	for i, card := range d {
		fmt.Println(i, card)
	}
}
