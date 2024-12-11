package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	
	// byteSlice := make([]byte, 999999)
	// res.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))

	lw := logWriter{}

	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Length: ", len(bs))
	return len(bs), nil
}