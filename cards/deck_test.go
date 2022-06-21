package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got %v", len(d))
	}
	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs but got %v", d[0])
	}
	if d[len(d)-1] != "King of Harts" {
		t.Errorf("Expected last card to be King of Harts but got %v", d[len(d)-1])

	}
}
