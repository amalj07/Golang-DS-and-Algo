package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) breadthFirstSearch() {
	nodes := []*Node{}

	nodes = append(nodes, bst.root)
	for len(nodes) > 0 {
		currentNode := nodes[0]
		fmt.Print(nodes[0].value, " ")
		nodes = nodes[1:]
		if currentNode.left != nil {
			nodes = append(nodes, currentNode.left)
		}

		if currentNode.right != nil {
			nodes = append(nodes, currentNode.right)
		}
	}
}

func main() {
	bst := BinarySearchTree{}

	bst.breadthFirstSearch()
}
