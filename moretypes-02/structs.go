package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// 変数にVertex型の値を入れて中身変えてログ出力
	v := Vertex{3, 4}
	v.X = 5
	v.Y = 6
	fmt.Println(v)
}
