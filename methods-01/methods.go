package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// vはレシーバ引数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 普通のやつ
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// v Vertex型のメソッドを呼び出す
	fmt.Println(v.Abs())
	// 普通に呼び出す
	fmt.Println(Abs2(v))
}
