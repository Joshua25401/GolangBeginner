package main

import (
	"fmt"
)

func main() {
	name := "Joshua Ryandafres"

	for i := 0; i < len(name); i += 2 {
		fmt.Printf("%c", name[i])

		if name[i] == 'a' && name[i+2] == 'R' {
			break
		}
	}
}
