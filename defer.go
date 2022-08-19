package main

import (
	"fmt"
)

func logging() {
	fmt.Println("LOG")
}

func runApp(num int) {
	defer logging()
	println(10 / num)
}

func main() {
	runApp(0)
}
