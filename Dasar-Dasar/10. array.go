package main

import "fmt"

// Tipe Data Array
func main() {
	var names [3]string
	names[0] = "Mahendra"
	names[1] = "Chevro"
	names[2] = "Susanto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat Array Secara Langsung
	var values = [3]int{
		90,
		95,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])

	fmt.Println(len(names))
	fmt.Println(len(values))

	// Membuat Array Tanpa Jumlah Elemen
	var numbers = [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(numbers)
}
