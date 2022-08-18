package main

import "fmt"

func main() {
	const firstName string = "john"
	const lastName = "mayer"
	const bin bool = false

	// this would be an error
	/*
	lastname = "doe" # cannot assign to lastName (untyped string constant "mayer")
	*/
	fmt.Printf("%s %s %t\n", firstName, lastName, bin)
}