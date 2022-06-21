package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveDeck("my_cards")
	hand, pile := deal(cards, 5)
	fmt.Println("----hand----")
	hand.print()
	fmt.Println("---pile---")
	pile.print()
}
