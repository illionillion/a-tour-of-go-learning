package main

import "fmt"

func main() {
	var i interface{}
	// var i any
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
// func describe(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
