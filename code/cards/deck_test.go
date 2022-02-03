package main

import (
	"os" // Package to access to system functionality
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Since deck is a type === extension of slice of strings --> len property can be used
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d)) // %v to add the next argument
		t.Errorf("Expected deck length of 16, but got %v", len(d)) // Another way to insert a value as an interpolation
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // Remove since in each execution this file will be created

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	myName := []string{"A", "L", "F", "R", "E", "D", "O"}
	deck := deck(myName)

	firstPart, secondPart := deal(deck, 3)

	if len(firstPart) != 3 && firstPart[0] != "A" && firstPart[1] != "L" && firstPart[2] != "F" {
		t.Errorf("Excpected 3 elements, but got #{len(firstPart)}, with values ALF, but got #{firstPart.toString()}")
	}

	if len(secondPart) != 4 && firstPart[3] != "R" && firstPart[4] != "E" && firstPart[5] != "D" && firstPart[6] != "O" {
		t.Errorf("Excpected 4 elements, but got #{len(secondPart)}, with values ALF, but got #{secondPart.toString()}")
	}
}
