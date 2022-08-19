package main

import (
	"fmt"
)

func getTwoValue() (int, string) {
	return 5, "Joe"
}

func getThreeValue() (int, byte, bool) {
	return 19, 'a', false
}

func namedGetValue() (a string, b string, c string) {
	a = "lol"
	b = "test"
	c = "pqkt"
	return c, b, a
}

func main() {
	a, b := getTwoValue()
	fmt.Println(a, b)

	f, _, e := getThreeValue()
	println(f, e)

	x, y, z := namedGetValue()
	println(x, y, z)
}
