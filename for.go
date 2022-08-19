package main

import (
	"fmt"
)

func main() {
	var counter int = 1
	for counter <= 12 {
		fmt.Println(counter)
		counter += 1
	}
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello World")
	}
	

	slice := []string{
		"Jonathan", "Joseph", "Jotaro", "Josuke", "Giorno",
	}
	mapp := make(map[string]string)
	mapp["a"] = "alpha"
	mapp["b"] = "beta"

	for idx, el := range(slice) {
		fmt.Printf("%d %s\n", idx, el)
	}

	for key, val := range(mapp) {
		fmt.Println(key, val)
	}
}
