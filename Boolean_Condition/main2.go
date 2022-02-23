package main

import (
	"fmt"
)

// Chained Conditional
// AND, OR, NOT
func main() {
	// Chain Condition
	cond := (true || false) && !false || !true
	// Output
	fmt.Println("Hasilnya adalah", cond)
}
