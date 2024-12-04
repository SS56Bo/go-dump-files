package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int64
	phone   int64
}

type personInfo struct {
	firstName string
	lastName  string
	gender    string
	contact   contactInfo
}

func main() {
	ss := personInfo{
		firstName: "Soham",
		lastName:  "Sarkar",
		gender:    "Male",
		contact: contactInfo{
			email:   "somethingDark@gmail.com",
			zipcode: 12345,
			phone:   12345,
		},
	}

	fmt.Println(ss)
}