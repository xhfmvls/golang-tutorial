package main

import "fmt"

func main() {
	var one int = 1
	var name string
	var boolVariable bool = false

	var (
		firstName = "vincent"
		lastName = "pradipta"
	)

	// another form of variable assignment
	test := "Test Variable"

	name = "Test"

	fmt.Println("ONE: ", one)
	fmt.Println("Name: ", name)
	fmt.Println("Char: ", name[1])
	fmt.Println("Len: ", len(name))
	fmt.Println("Inverse Bool", !boolVariable)
	fmt.Printf("Test: %s\n", test)
	fmt.Printf("Name: %s %s\n", firstName, lastName)
}