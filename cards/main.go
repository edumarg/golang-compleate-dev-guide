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
	hand, pile := deal(cards, 5)
	fmt.Println("----hand----")
	hand.print()
	fmt.Println("---pile---")
	pile.print()

}
