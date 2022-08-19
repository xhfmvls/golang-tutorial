package main

import (
	"fmt"
	"errors"
)

func div(a int, b int) (int, error) {
	if(b == 0) {
		return 0, errors.New("Number can't be devided by 0")
	} 
	return a / b, nil
}

func main() {
	num, err := div(30, 0)
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err.Error())
	}
}
