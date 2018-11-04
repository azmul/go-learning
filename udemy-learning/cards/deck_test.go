package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newCard()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func LastDeckTest(t *testing.T) {
	d := newCard()
	if d[len(d)-1] != "four of Clubs" {
		t.Errorf("Expected value of four of Clubs but got %v", d[len(d)-1])
	}
}
