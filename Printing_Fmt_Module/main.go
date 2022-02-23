package main

import "fmt"

/*
	FMT Cheat Sheet

	General
		%v	the value in a default format
		%+v	the value and the type
		%T 	the type of the value
		%%	a literal percent sign

	Boolean
		%t	the boolean value

	Integer
		%b	base 2
		%o	base 8
		%d	base 10
		%x	base 16
*/

func main() {
	nama := "Joshua Pangaribuan"
	var angka int = 100
	desimal := 200.33
	var boolean bool

	fmt.Printf("%T %v\n", nama, nama)
	fmt.Printf("%T %v\n", angka, angka)
	fmt.Printf("%T %v\n", desimal, desimal)
	fmt.Printf("%T %v\n", boolean, boolean)
}
