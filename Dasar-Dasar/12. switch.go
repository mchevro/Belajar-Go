package main

import "fmt"

func main() {
	name := "Mahendra"

	// Switch Dengan Kondisi
	switch name {
	case "Mahendra":
		fmt.Printf("Hii %s", name)
	case "Budi":
		fmt.Printf("Hii %s", name)
	default:
		fmt.Println("Siapa Nama Mu?")
	}

	// Switch Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("\nNama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// Swtich Tanpa Kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
