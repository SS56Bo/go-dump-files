package main

func main() {
	//cards := newDeck()
	cards := newDeckFromFile("my_cards")
	//result := cards.toString()
	//cards.toHardDrive("my_cards")

	cards.print()
}
