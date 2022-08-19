package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Println("Hello World!!!")
}

func main() {
	helloWorld()

	for i := 0; i < 10; i++ {
		helloWorld()
	}
}
