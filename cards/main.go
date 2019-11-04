package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("====== Hand ====")
	// hand.print()
	// fmt.Println("====== remaining cards ====")
	// remainingCards.print()
	// cards := newDeck()

	// cards.saveToFile("my-card.txt")

	cards := newDeckFromFile("my-deck.txt")
	cards.print()

	// cards := newDeck()

	// cards.shuffle()
	// cards.print()
}
