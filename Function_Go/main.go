package main

import (
	"fmt"
)

// Fibb Function
func fibb(angka int) int {
	if angka <= 1 {
		return 1
	}
	return fibb(angka-1) + fibb(angka-2)
}

// Hello World Function
func helloWorld() {
	fmt.Println("Hello World")
}

// printMhs Function
// parameter berupa slice string
func printMhs(listMhs []string) {
	fmt.Println("Index\t: Nama Mahasiswa")
	for index, mhs := range listMhs {
		fmt.Println(index, "\t:", mhs)
	}
}

// Return multiple values from function
func subsAndAdd(bil1, bil2 int) (substraction int, addition int) {
	substraction = bil1 - bil2
	addition = bil1 + bil2
	return
}

func main() {
	// Call Fibb Function
	angka := 10
	for i := 0; i < angka; i++ {
		fmt.Print(fibb(i), " ")
	}

	// Call Hello World Function
	fmt.Print("\n\n")
	helloWorld()

	// Call printMhs Function
	fmt.Println("")
	listMhs := []string{
		"Joshua Ryandafres",
		"Alex Sander",
		"Irwan Siagian",
	}
	printMhs(listMhs)

	// Call Multiple Return Function
	fmt.Println("")
	bil1, bil2 := 5, 3
	jawaban1, jawaban2 := subsAndAdd(bil1, bil2)
	fmt.Println("Jawaban 1:", jawaban1)
	fmt.Println("Jawaban 2:", jawaban2)
}
