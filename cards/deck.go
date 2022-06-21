package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func deal(d deck, handSize uint) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func stringToByte(s string) []byte {
	return []byte(s)

}

func (d deck) saveDeckToFile(fileName string) error {
	deckString := d.toString()
	deckByte := stringToByte(deckString)
	return ioutil.WriteFile(fileName, deckByte, 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Could not load deck from file. Error:", err)
		os.Exit(1)
	}
	fileString := string(bs)
	stringSlice := strings.Split(fileString, ",")
	return deck(stringSlice)
}
