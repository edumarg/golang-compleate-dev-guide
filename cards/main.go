package main

import "fmt"

func main() {
	cards := newDeck()
	hand, pile := deal(cards, 5)
	fmt.Println("----hand----")
	hand.print()
	fmt.Println("---pile---")
	pile.print()
}
