package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}

	// Syntax : arr[startIndex:endIndex]
	// Note : endIndex is not included
	slices := arr[1:5]
	fmt.Println(slices)
	// cap() : return the capacity of the slice or array
	slices = arr[:cap(arr)]
	fmt.Println(slices)

	// Slice of Mahasiswa
	var listMahasiswa []string = []string{
		"Joshua Pangaribuan",
		"Alex Sander",
		"Irwan Siagian",
	}
	// Output :
	fmt.Println(listMahasiswa)

	// Make Slice w/ function make()
	// Syntax : make([]type, length, capacity)
	// Note : capacity is optional
	var listMahasiswa2 = make([]string, 3)
	listMahasiswa2 = append(listMahasiswa2, "Joshua Pangaribuan")
	listMahasiswa2 = append(listMahasiswa2, "Alex Sander")
	for i := 0; i < len(listMahasiswa2); i++ {
		fmt.Printf("%d. %s\n", i, listMahasiswa2[i])
	}

}
