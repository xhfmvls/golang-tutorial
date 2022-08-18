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

	var newArr = [3]int{
		1, 2, 3,
	}

	fmt.Println(newArr)
	fmt.Println(newArr[1])

	names := [...]string{
		"vincent", "bloo", "john", "giorno", "leone", "joseph", "jonathan", "wamuu", "speedwagon", "ermes",
	}
	fmt.Println(names[3:6])
	fmt.Println(names[4:])
	fmt.Println(names[:5])
	fmt.Println(append(names[:5], "APPENDED"))
	fmt.Println(len(names[:5]))
	fmt.Println(cap(names[:5]))
}
