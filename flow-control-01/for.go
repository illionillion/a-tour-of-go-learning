package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	for j := 0; j < 10; j++ {
		sum += j
	}
	fmt.Println(sum)
}
