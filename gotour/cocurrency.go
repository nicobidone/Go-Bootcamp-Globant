package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

var x int

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
/*
func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
		x++
		if x == 10 {	close(ch) //Arreglo del range	}
	}
}
*/

func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)

}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	x = 0
	for i := range ch1 {
		fmt.Print("->", i)
	}
	fmt.Println()
	x = 0
	for i := range ch2 {
		fmt.Print("->", i)
	}
	return true
}

func main() {
	//ch := make(chan int)
	//go Walk(tree.New(1), ch)
	//for i := range ch {
	//	fmt.Println("->", i)
	//}
	Same(tree.New(1), tree.New(1))

}
