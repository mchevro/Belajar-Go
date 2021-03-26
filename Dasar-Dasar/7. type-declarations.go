package main

import "fmt"

// Type Declarations
func main() {
	type NoKTP string
	type Married bool
	type No int8

	var noKtpEko NoKTP = "123123123123"
	var marriedStatus Married = true
	var angka No = 127
	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
	fmt.Println(angka)
}
