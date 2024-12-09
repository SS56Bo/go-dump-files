package main

import "fmt"

func printMap(c map[string]string){
	for author, title := range c {
		fmt.Println("Author: ", author,",Title: ", title)
	}
}

func main() {
	//syntax = map[key data_type]value data_type{}
	books := map[string]string{
		"Harry Potter": "J.K Rowling",
		"A Tale of Two Cities": "Charles Dickens",
		"James Bond": "Ian Fleming",
	}

	destinations := map[string]string{
		"Paris": "France",
		"Sapienza": "Italy",
		"Marrakesh": "Morocco",
		"Bangkok": "Thailand",
		"Colorado": "USA",
		"Hokkaido": "Japan",
		"Hawke's Bay": "New Zealand",
		"Miami": "USA",
		"Santa Fortuna":"Columbia",
		"Mumbai": "India",
		"Whittleton Creek": "USA",
		"Isle of Skail": "North Atlantic Ocean",
		"New York": "USA",
		"Haven Island": "Maldives",
		"Dubai": "UAE",
		"Dartmoor": "England",
		"Berlin": "Germany",
		"Chonqking": "China",
		"Mendoza": "Argentina",
	}

	//var store map[string]int64

	sstore := make(map[string]int64)
	sstore["first"]=56
	sstore["second"]=41

	fmt.Println(destinations)
	//fmt.Println(store)
	fmt.Println(sstore)
	printMap(books)

}

