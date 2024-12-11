package main

import "fmt"

type bot interface {
	getMessage() string
}
//now this program has a new type called bot

func getGreeting(b bot) {
	fmt.Println(b.getMessage())
}

type englishBot struct{}
type spanishBot struct{}
type germanBot struct{}

func main(){
	eb := englishBot{}
	sb := spanishBot{}
	gb := germanBot{}

	getGreeting(eb)
	getGreeting(gb)
	getGreeting(sb)
}

func (englishBot) getMessage() string {
	return "Hello there!"
}

func (spanishBot) getMessage() string {
	return "Hola!"
}

func (germanBot) getMessage() string {
	return "Hallo!"
}