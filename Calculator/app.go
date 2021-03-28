package main

import (
	"fmt"
)

// ========================
// 		FUNGSI OPERASI
// ========================
func multiply(x, y int) int {
	return x * y
}

func sum(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func divide(x, y int) int {
	return x / y
}

// ========================
// 		RUN PROGRAM
// ========================
func main() {
	var (
		num1 int
		num2 int
		opsi int8
	)
	for {
		// Menu Operasi
		fmt.Println("[1] Penjumlahan")
		fmt.Println("[2] Perkalian")
		fmt.Println("[3] Pengurangan")
		fmt.Println("[4] Pembagian")
		fmt.Print("Pilih Angka Opsi : ")
		fmt.Scan(&opsi)

		// Input User
		fmt.Print("\nInput Angka Pertama : ")
		fmt.Scan(&num1)
		fmt.Print("Input Angka Kedua : ")
		fmt.Scan(&num2)

		// Seleksi Hasil Pemilihan Operasi

		switch opsi {
		case 1:
			count := sum(num1, num2)
			fmt.Printf("Result : %d\n\n", count)
		case 2:
			count := multiply(num1, num2)
			fmt.Printf("Result : %d\n\n", count)
		case 3:
			count := subtract(num1, num2)
			fmt.Printf("Result : %d\n\n", count)
		case 4:
			count := divide(num1, num2)
			fmt.Printf("Result : %d\n\n", count)
		default:
			fmt.Println("Opsi Tidak Ada")
		}
	}
}
