package main

import (
	"fmt"
	"time"
)

// ========================
// 		FUNGSI OPERASI
// ========================
func multiply(x, y int) int {
	var (
		result int
	)
	result = x * y
	return result
}

func sum(x, y int) int {
	var (
		result int
	)
	result = x + y
	return result
}

func subtract(x, y int) int {
	var (
		result int
	)
	result = x - y
	return result
}

func divide(x, y int) int {
	var (
		result int
	)
	result = x / y
	return result
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
	// Menu Operasi
	fmt.Println("[1] Penjumlahan")
	fmt.Println("[2] Perkalian")
	fmt.Println("[3] Pengurangan")
	fmt.Println("[4] Pembagian")
	fmt.Print("Pilih Angka Opsi : ")
	fmt.Scan(&opsi)

	// Input User
	fmt.Print("Input Angka Pertama : ")
	fmt.Scan(&num1)
	fmt.Print("Input Angka Kedua : ")
	fmt.Scan(&num2)

	// Seleksi Hasil Pemilihan Operasi
	if opsi == 1 {
		count := sum(num1, num2)
		fmt.Printf("Result : %d", count)
		time.Sleep(5 * time.Second)
	} else if opsi == 2 {
		count := multiply(num1, num2)
		fmt.Printf("Result : %d", count)
		time.Sleep(5 * time.Second)
	} else if opsi == 3 {
		count := subtract(num1, num2)
		fmt.Printf("Result : %d", count)
		time.Sleep(5 * time.Second)
	} else if opsi == 4 {
		count := divide(num1, num2)
		fmt.Printf("Result : %d", count)
		time.Sleep(5 * time.Second)
	}
}
