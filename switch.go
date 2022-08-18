package main

import (
	"fmt"
)

func main() {
	name := "Leone"
	// name := "Leon"
	// name := "leone"

	switch name {
	case "Eko":
		fmt.Println("Test")
	case "Leone":
		fmt.Println("OK")
	default:
		fmt.Println("Default")
	}

	switch length := len(name); length >= 5 {
	case true:
		fmt.Println("OK")
	case false:
		fmt.Println("NOT OK")
	}

	// num := 3
	num := 10

	switch {
	case num > 5:
		fmt.Println("BIG")
	default:
		fmt.Println("small")
	}
}
