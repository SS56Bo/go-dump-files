package main

import "fmt"

func main() {
	i := 1
	for i <= 8{
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
		i=i+1
	}
}