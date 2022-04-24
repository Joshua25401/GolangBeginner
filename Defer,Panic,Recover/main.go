package main

import (
	"fmt"
)

/*
1. defer keyword
	- digunakan untuk mengeksekusi function atau baris kode walaupun baris code sebelumnya error
	- biasanya digunakan untuk menutup koneksi seperti ke DB ( Database )

2. panic keyword
	- digunakan untuk menghentikan program
	- defer function keyword akan tetap di eksekusi
	- biasanya dipanggil untuk menunjukkan error pada program

3. recover keyword
	- adalah function yang digunakan untuk mendapatkan data panic
	- jadi ketika panic maka program tidak akan sepenuhnya menghentikan aplikasi

*/

func logging() {
	message := recover()
	fmt.Println("Error Message ", message)
}

func print_me() {
	defer logging()
	fmt.Println("Hello World!")
}

func print_us(value bool) {
	defer logging()

	if value == true {
		panic("ERROR")
	}
	fmt.Println("HELLO :)")
}

func main() {
	print_me()
	print_us(true)
	print_us(false)
}
