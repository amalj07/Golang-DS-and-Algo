package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (linkedList *LinkedList) Append(value int) {
	newNode := Node{}
	newNode.value = value
	newNode.next = nil

	if linkedList.length == 0 {
		linkedList.head = &newNode
		linkedList.length++
		return
	}

	ptr := linkedList.head

	for i := 0; i < linkedList.length; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			linkedList.length++
			return
		}
		ptr = ptr.next
	}
}

func main() {
	l := LinkedList{}
	l.Append(10)
	fmt.Print(l.head)
}
