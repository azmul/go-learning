package main

func main() {
	cards := newCard()
	// hand, remaingCards := deal(cards, 5)
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.shuffle()
	cards.print()
}
