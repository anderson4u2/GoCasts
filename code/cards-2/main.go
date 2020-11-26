// package main

// import "fmt"

// func main() {
// 	// var card string = "Ace of Spades"
// 	card := "Ace of Spades"
// 	card = "Five of Diamonds"
// 	fmt.Println(card)
// }

// package main

// import "fmt"

// var deckSize int

// func main() {
// 	deckSize = 50
// 	fmt.Println(deckSize)
// }

// 15. Function and Return Types

// package main

// import "fmt"

// func main() {
// 	card := newCard()

// 	fmt.Println(card)
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }

// 16. Slices and For Loops

// package main

// import "fmt"

// func main() {
// 	// slice
// 	cards := []string{"one", "two"}
// 	// It doesn't modify existing `cards`, it returns a new one, and assigns it to variable called `cards`
// 	cards = append(cards, "three")

// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
// 	// fmt.Println(cards)
// }

// 18. Custom Type Declarations

package main

func main() {
	// slice
	cards := deck{"one", "two"}
	// It doesn't modify existing `cards`, it returns a new one, and assigns it to variable called `cards`
	cards = append(cards, "three")

	cards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)
}
