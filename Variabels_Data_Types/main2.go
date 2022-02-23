package main

import "fmt"

func main() {
	// Implicit and Explicit Var Declaration

	// Implicit Declaration
	var number = 2
	fmt.Printf("%T\n", number)

	nama := "Joshua Pangaribuan"
	fmt.Printf("%T %v\n", nama, nama)

	// Default type
	// Ternary operator in golang
	var str string
	c := (map[bool]string{true: "Empty", false: "Not Empty"})[str == ""]
	fmt.Println(c)

}
