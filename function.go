package main

import (
	"fmt"
)

type F func(int, int) int

func execOperation(num1 int, num2 int, f F) int {
	return f(num1, num2)
}

func main() {
	fmt.Println(execOperation(10, 13, func(a int, b int)int {
		return a + b
	}))

	fmt.Println(execOperation(10, 13, func(a int, b int)int {
		return a - b
	}))
}
