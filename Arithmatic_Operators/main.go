package main

import (
	"fmt"
)

func main() {
	// Variabel
	var angka int = 10
	var angka2 float64 = 20.33

	// Typecasting
	hasil := angka + int(angka2)

	// Output
	fmt.Println("Hasilnya adalah", hasil)
}
