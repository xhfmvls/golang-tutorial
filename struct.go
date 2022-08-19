package main

import (
	"fmt"
)

type Student struct {
	Name, NIM string
	Age int

}

func(student Student) greetings(name string) {
	fmt.Printf("Hello %s, my name is %s, and i'm %d years old\n", name, student.Name, student.Age)
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
	stu1.greetings("Narancia")
	stu2.greetings("Fugo")
}
