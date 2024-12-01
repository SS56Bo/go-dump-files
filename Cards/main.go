package main

import "fmt"

func main() {
	cards := newDeck()

	hands, remainingCards := deal(cards, 5)
	hands.print()
	fmt.Println(" ")
	remainingCards.print()
}
