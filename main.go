package main

import "fmt"

func main() {

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 3)
	// fmt.Println("Cards in hand")
	// hand.print()
	// fmt.Println("Remain cards")
	// remainingCards.print()

	cards := newDeck()

	fmt.Println(cards.saveToFile("my_cards"))

}
