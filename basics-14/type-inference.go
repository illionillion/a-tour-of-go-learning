package main

import "fmt"

func main() {
	v := 42 // 変えると暗黙的に型も変わる
	fmt.Printf("v is of type %T\n", v)
}
