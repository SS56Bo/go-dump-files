package main

import "fmt"

type bot interface {
	getMessage() string
}

func getGreeting(b bot) {
	fmt.Println(b.getMessage())
}

type englishBot struct{}
type spanishBot struct{}

func main(){
	eb := englishBot{}
	sb := spanishBot{}

	getGreeting(eb)
	getGreeting(sb)
}

func (englishBot) getMessage() string {
	return "Hello there!"
}

func (spanishBot) getMessage() string {
	return "Hola!"
}