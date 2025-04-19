package main

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Bark() {
	println("Woof! My name is", d.Name)
}
func (d Dog) GetAge() int {
	return d.Age
}
func (d Dog) SetAge(age int) {
	d.Age = age
}
func (d Dog) GetName() string {
	return d.Name
}
func (d Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) String() string {
	return "Dog{" + d.Name + ", " + fmt.Sprint(d.Age) + "}"
}
func (d Dog) Speak() {
	println("Woof! My name is", d.Name)
}
func (d Dog) Eat() {
	println("Eating...")
}
func (d Dog) Sleep() {
	println("Sleeping...")
}
func (d Dog) Play() {
	println("Playing...")
}
func (d Dog) Fetch() {
	println("Fetching...")
}
func (d Dog) Sit() {
	println("Sitting...")
}
func (d Dog) RollOver() {
	println("Rolling over...")
}

func main() {
	dog := Dog{Name: "Buddy", Age: 3}
	dog.Bark()
	dog.SetAge(4)
	println("Dog's age:", dog.GetAge())
	dog.SetName("Max")
	println("Dog's name:", dog.GetName())
	println(dog.String())
	dog.Speak()
	dog.Eat()
	dog.Sleep()
	dog.Play()
	dog.Fetch()
	dog.Sit()
	dog.RollOver()

}

