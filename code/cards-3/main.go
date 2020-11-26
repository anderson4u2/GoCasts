package main

// func main() {
// 	cards := newDeck()

// 	// hand, remainingCards := deal(cards, 5)

// 	// hand.print()
// 	// remainingCards.print()
// 	// fmt.Println(cards.toString())
// 	cards.saveToFile("test")
// }

// 28. Error Handling, but actually finishing the newDeckFromFile function
// func main() {
// 	cards := newDeckFromFile("test")
// 	cards.print()
// }

// 29. Shuffling a Deck
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
