package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello"+"world")
	fmt.Println(1+1)
	//AND 
	fmt.Println(true && false) 	// 1 AND 0 => 0
	fmt.Println(true && true) 	// 1 AND 1 => 1
	fmt.Println(false && false) // 0 AND 0 => 0
	fmt.Println(false && true) 	// 0 AND 1 => 0
	fmt.Println(" ")
	//OR
	fmt.Println(false || false) // 0 AND 0 => 0 
	fmt.Println(false || true) 	// 0 AND 1 => 1
	fmt.Println(true || false) 	// 1 AND 0 => 1
	fmt.Println(true || true) 	// 1 AND 1 => 1
	fmt.Println(" ")

	fmt.Println("13.0/3.0= ", 13.0/3.0)
	fmt.Println("15.74/2.41= ",15.74/2.41)
	fmt.Println("PI: ", 22.0/7.0)
	
}