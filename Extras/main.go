package main

import "fmt"

func main() {
	//create a slice
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, check := range nums {
		if (check%2==0){
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")
		}
	}
}
