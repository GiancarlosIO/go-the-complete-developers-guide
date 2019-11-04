package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %s", string(len(cards)))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Fource of Clubs, but got %s", cards[len(cards)-1])
	}
}

func TestSsaveToDeckAndNewDeckFromFile(t *testing.T) {
	var deckFile = "_decktesting.txt"
	os.Remove("decks/" + deckFile)

	deck := newDeck()
	deck.saveToFile(deckFile)

	loadedDeck := newDeckFromFile(deckFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("decks/" + deckFile)
}
