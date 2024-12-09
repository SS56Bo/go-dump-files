package main

import "fmt"

type bot interface {
	getMessage() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eB := englishBot{}
	sB := spanishBot{}

	getMessage(eB)
	getMessage(sB)
}

func getMessage(bots) string {
	fmt.Println()
}
