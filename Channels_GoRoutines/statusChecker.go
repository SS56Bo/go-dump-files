package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://ghijik.com",
	}

	c := make(chan string)

	for _, link := range links {
		// we have to pass the channel to the function
		go checkLink(link, c)
	}

	
	// for i := 0; i<len(links); i++{
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	//wait for the channel to return something
	// for l := range c {
	// 	time.Sleep(5*time.Second)
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5*time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down or may not exist")
		c<-link
		return
	}
	fmt.Println(link, "is up!")
	c <- "yep, it's up"
}
