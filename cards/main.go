package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveDeckToFile("my_cards")
	myDeck := newDeckFromFile("my_cards")
	fmt.Println("---from file---")
	myDeck.print()
	cards.shuffle()
	fmt.Println("---shuffled deck---")
	cards.print()
	hand1, pile := deal(cards, 5)
	hand2, pile := deal(pile, 5)
	hand3, pile := deal(pile, 5)
	fmt.Println("----hands----")
	fmt.Println("--hand 1--")
	hand1.print()
	fmt.Println("--hand 2--")
	hand2.print()
	fmt.Println("--hand 3--")
	hand3.print()
	fmt.Println("---pile---")
	pile.print()

}
