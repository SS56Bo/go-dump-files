package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		//if the len is not equal to 16, that error will be notified to the test handler Testing
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//Random tests
	// if d[0] != "Ace of Spades" {
	// 	t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	// }

	// if d[len(d)-1] != "King of Diamonds" {
	// 	t.Errorf("Expected last card to be King of Diamonds, but got %v", d[len(d)-1])
	// }
}

func checkIOTest(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.toHardDrive("_decktesting")
	
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}