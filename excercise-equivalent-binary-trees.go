package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecur(t, ch)
	close(ch)
}

func walkRecur(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecur(t.Left, ch)
	ch <- t.Value
	walkRecur(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for n := range ch1 {
		if n != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(10), tree.New(9)))
}
