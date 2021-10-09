package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type CircularLinkedList struct {
	head   *Node
	length int
}

func (circularLinkedList *CircularLinkedList) append(value int) {
	newNode := Node{}
	newNode.value = value

	if circularLinkedList.head == nil {
		circularLinkedList.head = &newNode
		newNode.next = &newNode
		circularLinkedList.length++
		return
	}

	ptr := circularLinkedList.head
	for ptr.next != circularLinkedList.head {
		ptr = ptr.next
	}

	newNode.next = ptr.next
	ptr.next = &newNode
	circularLinkedList.length++
	return
}

func (circularLinkedList *CircularLinkedList) prepend(value int) {
	newNode := Node{}
	newNode.value = value

	if circularLinkedList.head == nil {
		circularLinkedList.head = &newNode
		newNode.next = &newNode
		circularLinkedList.length++
		return
	}

	ptr := circularLinkedList.head
	for ptr.next != circularLinkedList.head {
		ptr = ptr.next
	}
	ptr.next = &newNode
	newNode.next = circularLinkedList.head
	circularLinkedList.head = &newNode

	circularLinkedList.length++
	return
}

func (circularLinkedList *CircularLinkedList) insert(position, value int) {
	if position <= 0 || position > circularLinkedList.length {
		fmt.Println("Error: Invalid position")
		return
	}

	ptr := circularLinkedList.head
	for i := 0; i < position; i++ {
		if i == position-2 {
			newNode := Node{}
			newNode.value = value
			newNode.next = ptr.next
			ptr.next = &newNode
			circularLinkedList.length++
			return
		}

		ptr = ptr.next
	}
}

func (circularLinkedList *CircularLinkedList) removeFromBeginning() {
	if circularLinkedList.head == nil {
		fmt.Println("Error: Linked list is empty")
		return
	}

	ptr := circularLinkedList.head
	for ptr.next != circularLinkedList.head {
		ptr = ptr.next
	}

	circularLinkedList.head = circularLinkedList.head.next
	ptr.next = circularLinkedList.head
	circularLinkedList.length--
	return
}

func (circularLinkedList *CircularLinkedList) removeFromEnd() {
	if circularLinkedList.head == nil {
		fmt.Println("Error: Linked list is empty")
		return
	}

	ptr := circularLinkedList.head
	for i := 0; i < circularLinkedList.length-2; i++ {
		ptr = ptr.next
	}
	ptr.next = circularLinkedList.head
	circularLinkedList.length--
	return
}

func (circularLinkedList *CircularLinkedList) remove(position int) {
	if position <= 0 || position > circularLinkedList.length {
		fmt.Println("Error: Invalid position")
		return
	}

	if position == 1 {
		circularLinkedList.removeFromBeginning()
		return
	}

	if position == circularLinkedList.length {
		circularLinkedList.removeFromEnd()
		return
	}

	ptr := circularLinkedList.head
	for i := 0; i < position-1; i++ {
		if i == position-2 {
			ptr.next = ptr.next.next
			circularLinkedList.length--
			return
		}
		ptr = ptr.next
	}
}

func (circularLinkedList *CircularLinkedList) display() {
	if circularLinkedList.head == nil {
		fmt.Print("Linked list is empty")
		return
	}

	ptr := circularLinkedList.head
	for i := 0; i < circularLinkedList.length; i++ {
		fmt.Print(ptr.value, " ")
		ptr = ptr.next
	}

	return
}

func main() {
	cl := CircularLinkedList{}
	cl.append(3)
	cl.append(7)
	cl.prepend(1)
	cl.prepend(5)
	cl.insert(2, 9)
	cl.removeFromBeginning()
	cl.removeFromEnd()
	cl.remove(3)
	cl.display()
}
