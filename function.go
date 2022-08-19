package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Println("Hello World!!!")
}

func greet(name string, msg string) {
	fmt.Println(msg, name)
}

func main() {
	helloWorld()

	for i := 0; i < 5; i++ {
		helloWorld()
	}

	greet("vincent", "hello")
}
