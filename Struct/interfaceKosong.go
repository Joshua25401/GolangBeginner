package main

import (
	"fmt"
	"reflect"
)

func printMe(message interface{}) {
	fmt.Println("Hello : ", message, reflect.TypeOf(message))
}

func main() {
	printMe("Joshua Pangaribuan")
	printMe(true)
	printMe(100.0)
}
