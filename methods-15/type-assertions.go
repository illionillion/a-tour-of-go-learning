package main

import "fmt"

func main() {
	var i any = "hello"
	
	// any型を文字列に変換
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic // , okすればパニックしない
	// fmt.Println(f)
}
