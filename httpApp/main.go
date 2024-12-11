package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	
	byteSlice := make([]byte, 999999)
	res.Body.Read(byteSlice)
	fmt.Println(string(byteSlice))
}

