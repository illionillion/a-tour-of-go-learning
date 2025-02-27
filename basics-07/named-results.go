package main

import "fmt"

// 返り値の型だけでなく、変数名も書ける
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// xとyを省略できる
	return
}

func main() {
	fmt.Println(split(17))
}
