package main

import (
	"fmt"
)

func logging() {
	fmt.Println("LOG")
}

func runApp(error bool) {
	defer logging()
	if(error) {
		panic("[-] ERROR")
	}
}

func main() {
	runApp(true)
}
