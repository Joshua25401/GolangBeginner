package main

import (
	"fmt"
)

func main() {
	// Initiate Map Explicit
	var mapNames map[int]string = map[int]string{
		1: "Joshua Ryandafres",
		2: "Alex Sander",
		3: "Irwan Siagian",
	}

	// Output
	fmt.Println(mapNames)

	// Add new value to map
	mapNames[4] = "Natalia Ervina"

	// Output
	fmt.Println(mapNames)

	// Accessing value from map
	fmt.Println(mapNames[1])

	// Delete value from map
	delete(mapNames, 3)
	fmt.Println(mapNames)

	// Check if key exist
	if _, exist := mapNames[3]; exist {
		fmt.Println("Key 3 exist")
	} else {
		fmt.Println("Key 3 not exist")
	}
}
