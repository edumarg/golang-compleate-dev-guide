package main

import "fmt"

// Create a new type of "deck"
// which is a slice of strings

type deck []string

// receiver function: a function that receives a variable.
// func (<variable name> <variable type>) <function name> {logic}

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Clubs", "Spades", "Diamonds", "Harts"}
	cardsValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() { // receiver function that receives any variable type deck and use it inside the function logic
	for _, card := range d {
		fmt.Println(card)
	}
}
