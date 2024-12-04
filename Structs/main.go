package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (pointerToPerson *person) updateFirstName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastName (newLastName string) {
	(*pointerToPerson).lastName = newLastName
}

func (p person) printStuff() {
	fmt.Printf("%v", p)
}
 
func main() {
	alex := person{lastName: "Wayne", firstName: "Bruce"}
	// fmt.Println(alex.firstName)
	// fmt.Printf("%v", alex.lastName)
	alexPointer := &alex
	alexPointer.updateFirstName("Damian")
	alexPointer.updateLastName("Grayson")
	alex.printStuff()
}