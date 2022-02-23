package main

import (
	"fmt"
)

func main() {
	var a int = 11
	switch a {
	case 10, 11:
		fmt.Println("a is 10 or 11")
	case 20:
		fmt.Println("a is 20")
	default:
		fmt.Println("a is not 10 or 20")
	}
}
