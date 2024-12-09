// beginner-level exercises

// Struct Manipulation
// Create a struct Book with fields: Title, Author, Year.
// Write a function that accepts a slice of Book and
// returns all books published after a certain year.

package main

import "fmt"

type Book struct {
	Title string
	Author string
	Year int32 
}

func main() {
	books := []Book{
		{Title: "Casino Royale", Author: "Ian Fleming", Year: 1953},
		{Title: "Brave New World", Author: "Aldous Huxley", Year: 1932},
		{Title: "Dune", Author: "Frank Herbert", Year: 1965},
		{Title: "The Hunt for Red October", Author: "Tom Clancy", Year: 1984},
		{Title: "1984", Author: "Frank Herbert", Year: 1965},
		{Title: "The Three Body Problem", Author: "Cixin Liu", Year: 2006},
	}	
	fmt.Printf("%v \n", printPublishedBook(books))
}

func printPublishedBook(b []Book) []Book{
	var FilterBooks []Book
	for _, book := range b {
		if book.Year>1950 {
			FilterBooks = append(FilterBooks, book)
		}
	}
	return FilterBooks
}