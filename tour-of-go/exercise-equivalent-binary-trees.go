package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var dfs func(t *tree.Tree)
	dfs = func(t *tree.Tree) {
		if t == nil {
			return
		}

		dfs(t.Left)
		ch <- t.Value
		dfs(t.Right)
	}

	dfs(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			return true
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false
}
