package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{lastName: "Wayne", firstName: "Bruce"}
	fmt.Println(alex.firstName)
	fmt.Printf("%v", alex.lastName)
}