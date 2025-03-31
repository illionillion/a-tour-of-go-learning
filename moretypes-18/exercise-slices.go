package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// なんかよくわからんけどbase64で画像生成
	result := make([][]uint8, dy)
	for y := range dy {
		row := make([]uint8, dx)
		for x := range dx {
			row[x] = uint8(x * y) // Example pattern
		}
		result[y] = row
	}
	return result
}

func main() {
	pic.Show(Pic)
}
