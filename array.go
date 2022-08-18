package main

import (
	"fmt"
)

func main() {
	var arr [3]string

	arr[0] = "Testing"
	arr[1] = "Bluu"
	arr[2] = "Owww"
	
	fmt.Printf("%s %s %s\n", arr[0], arr[1], arr[2])

	var newArr = [3]int {
		1, 2, 3,
	}

	fmt.Println(newArr);
	fmt.Println(newArr[1]);
}