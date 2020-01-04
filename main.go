package main

import (
	"fmt"
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

func main() {

	t := tree.New(1)
	channel := make(chan int)
	go walk(t, channel)

	fmt.Printf("the tree: %v\n", t)

	for v := range channel {
		fmt.Printf("received from channel: (%v)\n", v)
	}
	fmt.Println("finished walking the tree")
}
