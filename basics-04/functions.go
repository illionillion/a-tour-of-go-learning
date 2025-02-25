package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func substruct(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(substruct(42, 13))
}
