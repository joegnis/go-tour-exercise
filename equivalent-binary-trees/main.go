/*
A function to check whether two binary trees store the same sequence is
quite complex in most languages. We'll use Go's concurrency and channels
to write a simple solution.

Solution: adding all elements from both trees into one channel, and
counting if there are two for each number
*/
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	go func() {
		Walk(t1, ch)
		Walk(t2, ch)
		close(ch)
	}()
	valueCount := make(map[int]int)
	for v := range ch {
		_, ok := valueCount[v]
		if ok {
			valueCount[v]--
		} else {
			valueCount[v] = 1
		}
	}
	for _, count := range valueCount {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
