package main

import (
	"fmt"
)

// create a new deck
// which is a slice of string

type deck []string

func newCard() deck {
	cards := deck{}

	cardSuits := []string{"spaeds", "diamonds", "hearts", "Clubs"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
