package main

func main() {
	cards := newCard()

	hand, remaingCards := deal(cards, 5)

	hand.print()
	remaingCards.print()
}
