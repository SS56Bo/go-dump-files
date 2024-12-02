package main

import (
	"fmt"
	"os"
	"strings"
)

//create a new type of 'deck'
type deck []string

func newDeck() deck{
	newCards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardNumber := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten","King", "Queen", "Jack"}

	for _, suits := range cardSuits {
		for _, num := range cardNumber {
			newCards = append(newCards, num+" of "+suits)
		}
	}

	return newCards
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int ) (deck, deck){
	 return d[:handSize], d[handSize:]
}

func  (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func  (d deck) toHardDrive(fileName string) error{
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(filenmae string) deck {
	bs, err := os.ReadFile(filenmae)
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	store :=  strings.Split(string(bs), ",")
	return deck(store)
}