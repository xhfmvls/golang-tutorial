package main

import (
	"fmt"
	"strconv"
)

type Add func(int, int) int

func addFunc(a int, b int) int {
	return a + b
}

func helloAdd(name string, newFunc Add, num1 int, num2 int) string {
	return "Hello: " + name + " ... " + strconv.Itoa(newFunc(num1, num2))
}

func main() {
	fmt.Println(helloAdd("Vincent", addFunc, 10, 13))
}
