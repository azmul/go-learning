package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// print all cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// create handsize
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// all cards join to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save to file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Read file
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// create shuffle
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
