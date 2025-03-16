package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// シャローコピー
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
