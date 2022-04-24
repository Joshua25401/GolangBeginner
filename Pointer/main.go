package main

import "fmt"

func changeByReference(value *int) {
	*value = 10
	fmt.Println("Value in Func =", *value)
}

func changeByValue(value int) {
	value = 10
	fmt.Println("Value in Func =", value)
}

func main() {
	value := 5
	fmt.Println("Value in Main =", value)
	changeByReference(&value)
	//changeByValue(value)
	fmt.Println("Value in Main =", value)
}
