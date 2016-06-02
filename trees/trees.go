package main

import "fmt"
import "golang.org/x/tour/tree"

// Walk the tree in-order
// (left subtree, node, right subtree)
// and write the values to the channel

func walk(t *tree.Tree, c chan int) {
	// base case
	if t == nil {

	}

	// recursive step
	if t.Left != nil {
		walk(t.Left, c)
	}

	c <- t.Value

	if t.Right != nil {
		walk(t.Right, c)
	}
}

func Walk(t *tree.Tree, c chan int) {
	walk(t, c)
	close(c)
}

func dump(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	var tree = tree.New(1)
	var out = make(chan int, 10)
	go Walk(tree, out)
	dump(out)
}
