package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "bark"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meow"
}

func main() {
	var a Animal
	a = Dog{}
	fmt.Println(a.Speak())

	a = Cat{}
	fmt.Println(a.Speak())
}
