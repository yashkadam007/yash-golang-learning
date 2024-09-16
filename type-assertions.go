package main

import "fmt"

func main() {
	var i interface{} = "hellow world"

	a := i.(string)
	fmt.Print(a)

	b, ok := i.(float32)

	fmt.Println(b, ok)
}
