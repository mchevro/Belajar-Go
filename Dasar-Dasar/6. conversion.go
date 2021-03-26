package main

import "fmt"

// Konversi Tipe Data
func main() {

	//Note : Harus disesuaikan dengan isi dari valuenya
	var nilai32 int32 = 127
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Mchevro"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
