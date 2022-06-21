package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveDeckToFile("my_cards")
	hand, pile := deal(cards, 5)
	fmt.Println("----hand----")
	hand.print()
	fmt.Println("---pile---")
	pile.print()
	myDeck := newDeckFromFile("my_cards")
	fmt.Println("---from file---")
	myDeck.print()

}
