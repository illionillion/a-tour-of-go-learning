package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 3の倍数になるまで無限ループ
	for {
		num := rand.Intn(100)       // 乱数を変数に格納
		fmt.Println("生成された値:", num) // 乱数の値を出力
		if num%3 != 0 {
			break // 3の倍数でなければループを抜ける
		}
		fmt.Println("ループ")
	}
}
