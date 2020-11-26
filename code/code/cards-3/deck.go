package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function with multiple arguments and multiple return values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely yquit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Converting byte slice to a string, and then to a []string (slice of strings)
	s := strings.Split(string(bs), ",") // Spades of Ace,Spades of Two,Spades of Three,Spades of Four.....
	d := deck(s)

	return d
}

// 29. Shuffling a Deck
// func (d deck) shuffle() {
// 	for i := range d {
// 		newPosition := rand.Intn(len(d) - 1)

// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }

// 30. Random Number Generation
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
