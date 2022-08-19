package main

import (
	"fmt"
)

func newMap(name string) map[string]any {
	if name == "" {
		return nil
	} else {
		return map[string]any {
			"name": name,
		}
	}
}

func main() {
	flea := newMap("")
	// flea := newMap("Flea")
	// flea["status"] = true
	if flea == nil {
		fmt.Println("MT")
	} else {
		fmt.Println(flea)
	}
}
