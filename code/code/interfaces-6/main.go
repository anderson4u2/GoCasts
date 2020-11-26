package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// New value of type struct, when it has no fields associated
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// since we're not using the values passed next to receiver function,
// we can ommit them. Golang won't complain if declared and unused
func (englishBot) getGreeting() string {
	// VERY custom logic for generating english greetings
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
