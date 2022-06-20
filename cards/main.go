package main

import "fmt"

func main() {
	var card1 string = "Ace of Spades" // creating a variable named card1 with type of string and value "ace of spades"
	card2 := "Ace of Diamonds"
	card3 := newCard()                                                      // creating a variable that will auto detect it's type
	cards := []string{"Five of Clubs", "King of Spades", "Queen of Hearts"} // creating a slice of strings
	cards = append(cards, "Ten of Diamonds")
	fmt.Println("my cards:", card1, ",", card2, ",", card3)

	for i, card := range cards { // for loop to iterate over cards slice
		fmt.Println(i, card)
	}
}

func newCard() string { // declaring a function with name and type of the return
	return "Ace of Clubs"
}
