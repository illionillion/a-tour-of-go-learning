package main

import "fmt"

func adder() func(int) int {
	// 結果がここに引き継がれる
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
