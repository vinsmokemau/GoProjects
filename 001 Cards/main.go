package main

import "fmt"

func main() {
	// var card string = "Five of Diamonds"
	// card := "Five of Diamonds" Only ':=' to init the variable
	// card = "Ace of Spades" Only '=' in the reassignments
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
