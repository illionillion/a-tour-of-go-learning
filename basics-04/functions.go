package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// 型が同じなら省略できる
func substruct(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(substruct(42, 13))
}
