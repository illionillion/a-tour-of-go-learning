package main

import (
	"fmt"
	"strings"
)

func main() {
	// 2次元配列
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := range board {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
