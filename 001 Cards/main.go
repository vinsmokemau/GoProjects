package main

import "fmt"

func main() {
	// var card string = "Five of Diamonds"
	// card := "Five of Diamonds" Only ':=' to init the variable
	// card = "Ace of Spades" Only '=' in the reassignments
	// card := newCard() Create a variable with a function
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Picas") // Create a new slice and assign to the variable

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
