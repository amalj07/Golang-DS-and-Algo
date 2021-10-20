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

func (bst *BinarySearchTree) inOrderTraversal(node *Node) {
	if node.left != nil {
		bst.inOrderTraversal(node.left)
	}
	fmt.Print(node.value, " ")
	if node.right != nil {
		bst.inOrderTraversal(node.right)
	}
}

func (bst *BinarySearchTree) preOrderTraversal(node *Node) {
	fmt.Print(node.value, " ")
	if node.left != nil {
		bst.preOrderTraversal(node.left)
	}
	if node.right != nil {
		bst.preOrderTraversal(node.right)
	}
}

func (bst *BinarySearchTree) postOrderTraversal(node *Node) {
	if node.left != nil {
		bst.postOrderTraversal(node.left)
	}
	if node.right != nil {
		bst.postOrderTraversal(node.right)
	}
	fmt.Print(node.value, " ")
}

func main() {
	bst := BinarySearchTree{}
	bst.inOrderTraversal(bst.root)
	fmt.Println()
	bst.preOrderTraversal(bst.root)
	fmt.Println()
	bst.postOrderTraversal(bst.root)
}
