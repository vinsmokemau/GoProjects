package main

func main() {
	// var card string = "Five of Diamonds"
	// card := "Five of Diamonds" Only ':=' to init the variable
	// card = "Ace of Spades" Only '=' in the reassignments
	// card := newCard() Create a variable with a function
	// cards := deck{"Ace of Spades", newCard()} Create a deck with static and function
	// cards = append(cards, "Six of Picas") Create a new slice and assign to the variable
	// cards := newDeck() Create a Deck with our function in 'deck.go'

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	// hand, reaminingCards := deal(cards, 5)
	// hand.print()
	// reaminingCards.print()
}
