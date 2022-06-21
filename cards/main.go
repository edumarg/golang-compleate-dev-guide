package main

func main() {
	cards := deck{"Five of Clubs", "King of Spades", "Queen of Hearts"}
	cards = append(cards, "Ten of Diamonds")
	cards.print()
}
