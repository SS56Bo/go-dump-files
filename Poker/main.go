package main

import "fmt"

func main() {
	fmt.Println("Poker Game")
	fmt.Println("***************************")
	fmt.Println("Enter your choice: ")
	fmt.Println("1. Draw Deck")
	fmt.Println("2. Shuffle Deck")
	var choice string
	fmt.Scanln(&choice)
}