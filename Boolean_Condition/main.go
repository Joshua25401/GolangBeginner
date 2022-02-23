package main

import (
	"fmt"
)

func main() {
	bil1 := 5
	bil2 := 6.5

	var cond bool = float64(bil1)+2.5 > bil2
	var cond2 bool = "Josh" == "Josh"

	fmt.Println("Hasilnya adalah", cond)
	fmt.Println("Hasilnya adalah", cond2)
}
