package main

import "fmt"

var b, c, d bool

var e = 127

const Pie = 3.1428

func main() {
	var a int
	fmt.Println(a, b, c, d, e)
	e = 128
	fmt.Print(e)
	fmt.Println(Pie)
	fmt.Print(Pie)
}
