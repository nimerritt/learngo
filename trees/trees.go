package main

import "fmt"
import "golang.org/x/tour/tree"

// Walk the tree in-order
// (left subtree, node, right subtree)
// and write the values to the channel

// pass in callback by value?
func walk(t *tree.Tree, onVisit func(int)) {
	// base case
	if t == nil {

	}

	// recursive step
	if t.Left != nil {
		walk(t.Left, onVisit)
	}

	onVisit(t.Value)

	if t.Right != nil {
		walk(t.Right, onVisit)
	}
}

func Walk(t *tree.Tree, c chan int) {
	walk(t, func(v int) { c <- v })
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
