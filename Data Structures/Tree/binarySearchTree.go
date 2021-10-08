package main

import "fmt"

type Node struct {
	value int
	right *Node
	left  *Node
}

type BinarySearchTree struct {
	root *Node
}

func (binarySearchTree *BinarySearchTree) insert(value int) {
	newNode := Node{}
	newNode.value = value

	if binarySearchTree.root == nil {
		binarySearchTree.root = &newNode
		return
	}

	currentNode := binarySearchTree.root
	for true {

		if value == currentNode.value {
			fmt.Println("Duplicate value")
			return
		}

		if value < currentNode.value {
			if currentNode.left == nil {
				currentNode.left = &newNode
				return
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = &newNode
				return
			}
			currentNode = currentNode.right
		}
	}
}

func (binarySearchTree *BinarySearchTree) search(value int) *Node {
	if binarySearchTree.root == nil {
		fmt.Println("Tree is emtpy")
		return nil
	}

	currentNode := binarySearchTree.root

	for currentNode != nil {
		if value == currentNode.value {
			return currentNode
		} else if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	return nil
}

func main() {
	bst := BinarySearchTree{}

	bst.insert(53)
	bst.insert(7)
	bst.insert(21)
	bst.insert(3)
	bst.insert(76)
	bst.insert(2)
	bst.insert(43)
	bst.insert(87)
	bst.insert(39)
	fmt.Println(bst.search(43))
}
