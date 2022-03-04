package main

import "fmt"

func main() {
	// menggunakan syntax
	// var var_name var_type = value
	var nama string = "Joshua Pangaribuan"
	var nama2 string
	nama2 = "Irwan Siagian"
	fmt.Println(nama)
	fmt.Println(nama2)

	// Integers
	/*
		Unsigned integers
		uint8 -> capacity 0 - 255
		uint16 -> capacity 0 - 65535
		uint32 -> capacity 0 - 4294967295
		uint64 -> capacity 0 - 18446744073709551615

		Signed Integers
		int8 -> capacity -128 - 127
		int16 -> capacity -32768 - 32767
		int32 -> capacity -2147483648 - 2147483647
		int64 -> capacity -9223372036854775808 - 9223372036854775807
	*/
	var uint_1 uint8 = 255
	fmt.Println("Angka uint8 : ", uint_1)

	// Float
	/*
		float32 -> capacity -3.4028234663852886e+38 - 3.4028234663852886e+38
		float64 -> capacity -1.7976931348623157e+308 - 1.7976931348623157e+308
	*/
	var float_1 float32 = 3.14
	fmt.Println("Angka float32 : ", float_1)

	// Complex
	/*
		complex64 -> capacity -1.8446744073709552e+19 - 1.8446744073709552e+19
		complex128 -> capacity -1.8446744073709552e+19 - 1.8446744073709552e+19
	*/
	var complex_1 complex64 = 3 + 4i
	fmt.Println("Angka complex64 : ", complex_1)

	// Boolean
	/*
		bool -> capacity true - false
	*/
	var bool_1 bool = true
	fmt.Println("Angka bool : ", bool_1)

	var bool_2 bool
	fmt.Println("Angka bool : ", bool_2)
}
