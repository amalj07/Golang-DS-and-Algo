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
		linkedList.length++
		return
	}

	newNode.next = linkedList.head
	linkedList.head = &newNode
	linkedList.length++
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

// Add value to specified index of Linked list
func (linkedList *LinkedList) insert(index, value int) {
	newNode := Node{}
	newNode.value = value

	if index >= linkedList.length {
		linkedList.append(value)
		return
	}

	if index == 0 {
		linkedList.prepend(value)
		return
	}

	ptr := linkedList.head
	for i := 0; i != index-1; {
		ptr = ptr.next
		i++
	}
	newNode.next = ptr.next
	ptr.next = &newNode
	linkedList.length++
}

func (linkedList *LinkedList) display() {
	ptr := linkedList.head
	for i := 0; i < linkedList.length; i++ {
		fmt.Print(" ", ptr.value)
		ptr = ptr.next
	}
}

func main() {
	l := LinkedList{}
	l.append(10)
	l.append(15)
	l.prepend(4)
	l.insert(1, 12)
	l.prepend(23)
	l.insert(0, 1)
	l.insert(43, 31)
	l.append(21)
	l.display()
}
