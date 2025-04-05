package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func compute_greeting(fn func(string) string) string {
	return fn("Hello")

}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// 挨拶を生成するコールバック関数
	greetWithName := func(name string) string {
		return "Hello, " + name + "!"
	}

	greetFormally := func(name string) string {
		return "Good day, " + name + "."
	}

	// コールバック関数を使った挨拶
	fmt.Println(compute_greeting(greetWithName)) // "Hello, Hello!"
	fmt.Println(compute_greeting(greetFormally)) // "Good day, Hello."
}
