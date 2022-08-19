package main

import (
	"fmt"
	"reflect"
)

func printAny(a interface{}) {
	fmt.Println("Var:", a)
}

func emptyInterfaceReturn(a interface{}) interface{} {
	return reflect.TypeOf(a)
}

func main() {
	printAny(10)
	printAny("Test")
	printAny('T')
	printAny(false)
	fmt.Println(emptyInterfaceReturn(10))
	fmt.Println(emptyInterfaceReturn("Fak"))
	fmt.Println(emptyInterfaceReturn(emptyInterfaceReturn("Fak")))
	fmt.Println(emptyInterfaceReturn(false))
}
