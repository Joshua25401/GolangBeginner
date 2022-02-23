package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Initiate scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Ask user to enter the number of elements
	fmt.Print("Enter the number of elements : ")
	scanner.Scan()
	arrSize, err := strconv.ParseInt(scanner.Text(), 10, 64)

	// Check if the input is a valid number of elements
	if err == nil {
		arrList := make([]string, arrSize)

		// Ask user to enter the elements
		for i := 0; i < cap(arrList); i++ {
			fmt.Print("Enter the name : ")
			scanner.Scan()
			arrList[i] = scanner.Text()
		}

		// Print the output
		fmt.Printf("Index\t: Name\n")
		for index, name := range arrList {
			fmt.Printf("%d\t: %s\n", index, name)
		}
	} else {
		fmt.Println("Please enter a valid number! (", err, ")")
	}

}
