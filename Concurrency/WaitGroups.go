package main

import (
	"fmt"
)

func main() {

	nama := "joshua"

	for i := 0; i < len(nama); i++ {
		fmt.Printf("%03d", nama[i])
	}
}
