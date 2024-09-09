package main

import (
	"fmt"
	"math"
	"math/rand"
)

type MyInt int64

func (m MyInt) Abs() int64 {
	if m < 0 {
		return int64(-m)
	} else {
		return int64(m)
	}
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	v.Scale(10)
	fmt.Println(v.Abs())
	a := MyInt(-rand.Intn(20))
	fmt.Println(a.Abs())
}
