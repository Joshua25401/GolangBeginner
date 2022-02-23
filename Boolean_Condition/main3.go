package main

import (
	"fmt"
)

// If Else statement
func main() {
	name := "Josh"

	fmt.Println("Statement diluar If")
	cond := name == "josh" && name[0] != 'J'
	if !cond {
		fmt.Println("Hasil check kondisi : ", cond)
	}
	fmt.Println("Statement diluar If")
}
