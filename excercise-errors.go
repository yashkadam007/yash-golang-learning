package main

import (
	"fmt"
	"math"
)

type ErrNagativeSqrt float64

func (e ErrNagativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNagativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
