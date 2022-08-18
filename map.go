package main

import (
	"fmt"
)

func main() {
	// Alternative:
	// var person map[string]string = map[string]string{}

	// person := map[string: key]string: value{
	person := map[string]string{
		"name": "joseph joestar",
		"country": "Japan",
	}

	fmt.Println(person)
	fmt.Println(person["country"])

	person["major"] = "sociology"

	fmt.Println(person)

	person["country"] = "USA"

	fmt.Println(person)

	delete(person, "country")

	fmt.Println(person)

	// without data initialized
	book := make(map[string]string)
	book["test"] = "testing"
	fmt.Println(book)

	arrayOfBook := make([]string, 10)
	arrayOfBook[0] = "OK"
	fmt.Println(arrayOfBook)
	fmt.Println(cap(arrayOfBook))
	fmt.Println(len(arrayOfBook))
}
