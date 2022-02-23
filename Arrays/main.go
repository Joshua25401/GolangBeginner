package main

import (
	"fmt"
)

func main() {
	// Initiate Arrays
	var listMahasiswa [3]string = [3]string{"Joshua Ryandafres", "Irwan Siagian", "Alex Sander"}
	var daftarAngka [5]int

	// Output
	for i := 0; i < len(listMahasiswa); i++ {
		fmt.Println(listMahasiswa[i])
	}
	fmt.Println("Isi array : ", daftarAngka)
}
