package main

import "fmt"

func main() {
	/* For perulangan. mirip dengan while
	counter := 1
	for counter <= 10 {
		fmt.Println("Angka Ke", counter)
		counter++
	}
	*/

	//For dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Angka Ke", counter)
	}

	// For range
	slice := []string{"Mahendra", "Chevro", "Susanto"}
	for _, value := range slice { // Note : Simbol _ digunakan agar memberitahu golang bahwa variable tidak dibutuhkan
		fmt.Println(value)
	}
}
