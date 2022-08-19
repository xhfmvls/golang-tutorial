package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		if(i % 5 == 0) {
			continue
		}
		fmt.Println(i)
	}

	for i := 1; i <= 20; i++ {
		if(i % 5 == 0) {
			break
		}
		fmt.Println(i)
	}
}
