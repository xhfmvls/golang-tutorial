package main

import (
	"fmt"
	"strconv"
)

func main() {
	// between integer
	var value32 int32 = 10
	var valueIn64 int64 = int64(value32)

	fmt.Printf("%d %d\n", value32 + 7, valueIn64 + 5)
	
	// Int to String
	var stringOfInt string = strconv.Itoa(int(value32))
	fmt.Println(stringOfInt + "s")

	// Byte to String
	var name string
	name = "Leone Abbacchio"
	var fourthChar byte = name[5]
	fmt.Printf("%s %s\n", name, string(fourthChar))
}