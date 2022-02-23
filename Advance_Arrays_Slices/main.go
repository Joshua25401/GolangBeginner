package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Loop through the element
	// Using foreach
	names := []string{
		"Joshua Ryandafres",
		"Alex Sander",
		"Irwan Siagian",
	}

	// Foreach loop
	fmt.Printf("Index\t: Name\n")
	for index, name := range names {
		fmt.Printf("%d\t: %s\n", index, name)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name : ")
	scanner.Scan()
	names = append(names, scanner.Text())

	// Foreach loop
	fmt.Printf("New Element Added!\n")
	fmt.Printf("Index\t: Name\n")
	for index, name := range names {
		fmt.Printf("%d\t: %s\n", index, name)
	}
}
