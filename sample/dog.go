package main

import (
	"fmt"
)

// Animalインターフェース：外部に公開されたメソッドだけ
type Animal interface {
	Call()
	GetName() string
	GetAge() int
}

// Dog構造体：name, age は外から触れないように小文字
type Dog struct {
	name string
	age  int
}

// コンストラクタ関数（外部からdogを生成する唯一の方法）
func NewDog(name string, age int) Animal {
	return &Dog{name: name, age: age}
}

// インターフェースを満たすメソッド群
func (d *Dog) Call() {
	fmt.Printf("Woof! I'm %s (%d years old)\n", d.name, d.age)
}

func (d *Dog) GetName() string {
	return d.name
}

func (d *Dog) GetAge() int {
	return d.age
}

func main() {
	// Animal型として使用 → 外部からはdogのフィールドに触れない
	myDog := NewDog("Pochi", 3)
	myDog.Call()

	// fmt.Println(myDog.name) // コンパイルエラー：nameは非公開
}
