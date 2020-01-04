package binarytree

import (
	"golang.org/x/tour/tree"
)

func walk(node *tree.Tree, ch chan int) {
	walkInternal(node, ch)
	close(ch)
}

func walkInternal(node *tree.Tree, ch chan int) {
	if node.Left != nil {
		walkInternal(node.Left, ch)
	}

	ch <- node.Value

	if node.Right != nil {
		walkInternal(node.Right, ch)
	}
}

func same(t1, t2 *tree.Tree) bool {
	t1Ch, t2Ch := make(chan int), make(chan int)
	t1Vals, t2Vals := make([]int, 0), make([]int, 0)

	go walk(t1, t1Ch)
	go walk(t2, t2Ch)

	for t1v := range t1Ch {
		t1Vals = append(t1Vals, t1v)
	}

	for t2v := range t2Ch {
		t2Vals = append(t2Vals, t2v)
	}

	for i := 0; i < len(t1Vals); i++ {
		if t1Vals[i] != t2Vals[i] {
			return false
		}
	}

	return true
}
