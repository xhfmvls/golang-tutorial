package main

import (
	"fmt"
)

func main() {
	num := 12
	if num == 10 {
		fmt.Println(true)
	} else if num == 12 {
		fmt.Println("it's 12")
	} else {
		fmt.Println(false)
	}

	word := "it's all gonna be all right"
	if len(word) > 222 {
		fmt.Println("TOO LONG")
	} else {
		fmt.Println("Not Too Long")
	}
}
