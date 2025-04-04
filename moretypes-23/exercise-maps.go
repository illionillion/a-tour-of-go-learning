package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s) // 文字列をスペースで分割
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++ // 単語の出現回数をカウント
	}
	return counts
}

// IsOdd は与えられた整数が奇数かどうかを判定します。
func IsOdd(n int) bool {
	return n%2 != 0
}

func main() {
	wc.Test(WordCount)
}
