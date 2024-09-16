package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is of type int", v)
	case string:
		fmt.Println("This is of type string", v)
	default:
		fmt.Println("This is of god knows what type", v)
	}
}

func main() {
	do(23)
	do("hello")
	do(false)
}
