package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Add a new value to end of list
func (doublyLinkedList *DoublyLinkedList) append(value int) {
	newNode := Node{}
	newNode.value = value

	if doublyLinkedList.length == 0 {
		doublyLinkedList.head = &newNode
		doublyLinkedList.tail = &newNode
		doublyLinkedList.length++
		return
	}

	newNode.prev = doublyLinkedList.tail
	doublyLinkedList.tail.next = &newNode
	doublyLinkedList.tail = &newNode
	doublyLinkedList.length++
	return
}

// Add a value to beginning of list
func (doublyLinkedList *DoublyLinkedList) prepend(value int) {
	newNode := Node{}
	newNode.value = value
	if doublyLinkedList.length == 0 {
		doublyLinkedList.head = &newNode
		doublyLinkedList.tail = &newNode
		doublyLinkedList.length++
		return
	}

	newNode.next = doublyLinkedList.head
	doublyLinkedList.head.prev = &newNode
	doublyLinkedList.head = &newNode
	doublyLinkedList.length++
	return
}

func (doublyLinkedList *DoublyLinkedList) insert(position, value int) {
	if position < 0 {
		fmt.Println("Error: Invalid position")
		return
	}

	if position == 0 || position == 1 {
		doublyLinkedList.prepend(value)
		return
	}

	if position > doublyLinkedList.length {
		doublyLinkedList.append(value)
		return
	}

	ptr := doublyLinkedList.head
	for i := 0; i < position; i++ {
		if i == position-1 {
			newNode := Node{}
			newNode.value = value
			newNode.next = ptr
			newNode.prev = ptr.prev
			ptr.prev.next = &newNode
			ptr.prev = &newNode
			doublyLinkedList.length++
			return
		}
		ptr = ptr.next
	}
}

func (doublyLinkedList *DoublyLinkedList) removeFromBeginning() {
	if doublyLinkedList.length == 0 {
		fmt.Println("Linked list is emtpy")
		return
	}

	doublyLinkedList.head = doublyLinkedList.head.next
	doublyLinkedList.length--
	return
}

func (doublyLinkedList *DoublyLinkedList) removeFromEnd() {
	if doublyLinkedList.length == 0 {
		fmt.Println("Linked list empty")
		return
	}

	doublyLinkedList.tail = doublyLinkedList.tail.prev
	doublyLinkedList.tail.next = nil
	doublyLinkedList.length--
	return
}

func (doublyLinkedList *DoublyLinkedList) remove(position int) {
	if position <= 0 || position > doublyLinkedList.length {
		fmt.Println("Error: Invalid position")
		return
	}

	if position == 1 {
		doublyLinkedList.removeFromBeginning()
		return
	}

	if position == doublyLinkedList.length {
		doublyLinkedList.removeFromEnd()
		return
	}

	ptr := doublyLinkedList.head
	for i := 0; i < position; i++ {
		if i == position-1 {
			ptr.prev.next = ptr.next
			ptr.next.prev = ptr.prev
			doublyLinkedList.length--
		}
		ptr = ptr.next
	}
	return
}

func (doublyLinkedList *DoublyLinkedList) print() {

	if doublyLinkedList.length == 0 {
		fmt.Println("Linked list is empty")
		return
	}

	ptr := doublyLinkedList.head
	for i := 0; i < doublyLinkedList.length; i++ {
		fmt.Print(ptr.value, " ")
		ptr = ptr.next
	}
}
func main() {
	dll := DoublyLinkedList{}
	dll.append(4)
	dll.append(3)
	dll.prepend(1)
	dll.prepend(2)
	dll.insert(3, 8)
	dll.insert(1, 0)
	dll.remove(4)
	dll.print()
}
