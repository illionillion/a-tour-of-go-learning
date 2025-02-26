package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func calc(x []int) (int, int) {
	if len(x) == 0 {
		return 0, 0
	}
	max, min := x[0], x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func main() {
	a, b := swap("hello", "world")
	max, min := calc([]int{10, 12, 3, 4, 5})
	fmt.Println(a, b)
	fmt.Println("max is", max)
	fmt.Println("min is", min)
}
