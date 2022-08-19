package main

import (
	"fmt"
)

func fibo(n int)int {
	if n == 1 {
		return 1
	}
	return n * fibo(n - 1)
}

func main() {
	fmt.Println(fibo(5))

	val := fibo(6)
	fmt.Println(val)
}
