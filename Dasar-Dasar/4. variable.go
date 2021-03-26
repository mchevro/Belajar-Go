package main

import "fmt"

// Penulisan Variable
func main() {
	var name string

	name = "Mahendra Chevro"
	fmt.Println(name)

	name = "Mahendra Chevro Susanto"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// Langsung Deklarasikan Tanpa Memberitahun Tipe Datanya
	country := "Indonesia"
	fmt.Println(country)

	// Multipe Variable
	var (
		firstName string
		lastName  string
		oldAge    int8
	)
	firstName = "Eko"
	lastName = "Julyano"
	oldAge = 23
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(oldAge)
}
