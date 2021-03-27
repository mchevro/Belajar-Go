package main

import "fmt"


func main(){
	name := "Mahendra"
	
	switch name{
	case "Mahendra":
		fmt.Printf("Hii %s", name)
	case "Budi":
		fmt.Printf("Hii %s", name)
	default:
		fmt.Println("Siapa Nama Mu?")
	}
	
	switch length := len(name); length > 5{
	case true:
		fmt.Println("\nNama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
}