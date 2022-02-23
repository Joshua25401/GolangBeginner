package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Scan input dari user dan masukkan ke variabel userInput
	fmt.Printf("Ketikkan apa saja : ")
	scanner.Scan()
	userInput := scanner.Text()

	// Tampilkan output
	fmt.Printf("Input dari user : %q", userInput)

	// Conversi ke variabel bilangan
	fmt.Printf("\nKetikkan angka : ")
	scanner.Scan()
	bilangan, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Input dari user : \"%d\"", bilangan)

}
