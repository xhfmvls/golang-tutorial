package main

import (
	"fmt"
)

func main() {
	var(
		ta int = 5
		tb int = 10
	)
	var(
		eq1 bool = ta == tb
		eq2 bool = ta != tb
		eq3 bool = ta > tb 
		eq4 bool = ta < tb 
	)
	fmt.Printf("%t %t %t %t", eq1, eq2, eq3, eq4)
}