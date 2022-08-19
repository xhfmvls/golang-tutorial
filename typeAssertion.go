package main

import (
	"fmt"
)

func random() any {
	// return true
	return 12
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch result.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	default:
		fmt.Println("Else")
	}
}
