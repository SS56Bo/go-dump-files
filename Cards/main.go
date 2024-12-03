package main

func main() {
	cards := newDeck()
	//cards := newDeckFromFile("my_cards")
	//result := cards.toString()
	//cards.toHardDrive("my_cards")
	cards.shuffleDeck()
	cards.print()
	//fmt.Println(" ")
}
