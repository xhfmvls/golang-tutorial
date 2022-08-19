package main

import (
	"fmt"
)

func logging() {
	fmt.Println("LOG")
	errorMessage := recover()
	if(errorMessage == nil) {
		errorMessage = "No Error"
	}
	fmt.Println(errorMessage)
}

func runApp(error bool) {
	defer logging()
	if(error) {
		panic("[-] ERROR")
	}
}

func main() {
	runApp(true)
	runApp(false)
	fmt.Println("END")
}
