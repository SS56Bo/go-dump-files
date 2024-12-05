package main

import "fmt"

func main() {
	//syntax = map[key data_type]value data_type{}
	books := map[string]string{
		"Harry Potter":         "J.K Rowling",
		"A Tale of Two Cities": "Charles Dickens",
		"James Bond": "Ian Fleming",
	}

	fmt.Println(books)
}