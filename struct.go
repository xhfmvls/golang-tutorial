package main

import (
	"fmt"
)

type Student struct {
	Name, NIM string
	Age int

}

func main() {
	var stu1 Student
	stu1.Name = "Giorno"
	stu1.NIM = "0192381047"
	stu1.Age = 15
	fmt.Println(stu1)
	fmt.Println(stu1.NIM)

	// urutannya bebas
	stu2 := Student {
		Age: 22,
		NIM: "0183928172",
		Name: "Bruno",
	}
	fmt.Println(stu2)
}
