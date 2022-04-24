package main

import (
	"errors"
	"fmt"
)

/*
	Error Handling
		- Error handling di Go biasanya menggunakan package errors
		- Kemudian function harus mengembalikan nilai error
		- Untuk membuat nilai error baru menggunakan error.New(PesanError)

*/

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	total := nilai / pembagi
	return total, nil
}

func main() {
	hasil, error := pembagian(10, 5)

	if error != nil {
		fmt.Println("Hasil : ", hasil)
		fmt.Println("Error : ", error)
		return
	}
	fmt.Println("Hasil :", hasil)
}
