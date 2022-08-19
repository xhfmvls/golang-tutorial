package main

import (
	"fmt"
)

func multipyAll(numbers ...int) int {
	total := 1
	for _, val:= range(numbers) {
		total *= val
	}
	return total
}


func main() {
	fmt.Println(multipyAll(2, 3, 5, 6))
	numList := []int{ 
		1, 2, 3, 4,
	}
	fmt.Println(multipyAll(numList...))
}
