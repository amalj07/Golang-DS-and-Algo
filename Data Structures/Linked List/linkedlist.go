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

// Add value to begining of the list
func (linkedList *LinkedList) prepend(value int) {
	newNode := Node{}
	newNode.value = value

	if linkedList.length == 0 {
		newNode.next = nil
		return
	}

	newNode.next = linkedList.head
	linkedList.head = &newNode
	return
}

// Add a value to end of the list
func (linkedList *LinkedList) append(value int) {
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
	l.append(10)
	l.append(15)
	l.prepend(4)
	fmt.Print(l.head)
}
