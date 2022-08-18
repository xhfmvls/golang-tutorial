package main

import (
	"fmt"
)

func main() {
	var temp int = 10
	fmt.Printf("%d %.3f %d %d\n", temp + 3, float32(temp) / 4, temp - 7, temp * 2)

	var i int = 10
	i += 10
	fmt.Println(i)

	i--
	fmt.Println(i)

	var negative int = -10
	fmt.Println(negative)
}