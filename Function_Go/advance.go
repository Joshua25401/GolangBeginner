package main

import (
	"fmt"
)

// Normal Function w/ Parameter
func tambah(bil1, bil2 int) int {
	return bil1 + bil2
}

// Multiple return Function
func tambahKurang(bil1, bil2 int) (hasilTambah int, hasilKurang int) {
	hasilTambah = bil1 + bil2
	hasilKurang = bil1 - bil2
	return
}

func main() {
	// Function in variable (I)
	call := tambah
	fmt.Printf("Hasil : %d\n", call(2, 3))

	// Function in varible (II)
	functionVar := func(bil1, bil2 int) int {
		return bil1 - bil2
	}

	fmt.Printf("Hasil : %d\n", functionVar(8, 2))

	hasilTambah, hasilKurang := tambahKurang(10, 2)
	fmt.Printf("Hasil Tambah : %d\nHasil Kurang : %d\n", hasilTambah, hasilKurang)
}
