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

func (binarySearchTree *BinarySearchTree) delete(value int) {
	if binarySearchTree.root == nil {
		fmt.Println("Tree is emtpy")
		return
	}

	currentNode := binarySearchTree.root
	var parentNode *Node
	for currentNode != nil {
		// Find the node to delete
		if value < currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if value > currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.right
		} else { // Delete the node (value == currentNode.value)
			if currentNode.left == nil && currentNode.right == nil {
				// Node has zero child
				if parentNode.value > currentNode.value {
					parentNode.left = nil
					return
				} else {
					parentNode.right = nil
					return
				}
			} else if currentNode.right != nil || currentNode.left != nil {
				// Node has one child
				if parentNode.value > currentNode.value {
					if currentNode.left != nil {
						parentNode.left = currentNode.left
						return
					} else {
						parentNode.left = currentNode.right
						return
					}
				} else {
					if currentNode.left != nil {
						parentNode.right = currentNode.left
						return
					} else {
						parentNode.right = currentNode.right
						return
					}
				}
			} else {
				// Node has two children
				nodeToDelete := currentNode
				parentNode = currentNode
				// Find the largest node in the subtree
				for currentNode.right != nil {
					parentNode = currentNode
					currentNode = currentNode.right
				}

				nodeToDelete.value = currentNode.value
				parentNode.right = currentNode.left
				return
			}
		}
	}
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
	bst.insert(4)
	fmt.Println(bst.search(53))
	fmt.Println(bst.search(21))
	bst.delete(7)
	fmt.Println(bst.search(53))
}
