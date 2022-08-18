package main

import (
	"fmt"
)

func main() {
	type NoKtp string 
	type married bool
	var myKtpNumber NoKtp = "123212321"
	var areMarried married = false
	
	fmt.Println(myKtpNumber)
	fmt.Println(areMarried)
}