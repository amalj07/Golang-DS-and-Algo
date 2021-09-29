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

// Add a value to end of the list
func (doublyLinkedlist *DoublyLinkedList) append(value int) {
	newNode := Node{}
	newNode.value = value
	if doublyLinkedlist.length == 0 {
		doublyLinkedlist.head = &newNode
		doublyLinkedlist.tail = &newNode
		doublyLinkedlist.length++
		return
	}

	newNode.next = doublyLinkedlist.head
	newNode.prev = doublyLinkedlist.tail
	doublyLinkedlist.tail.next = &newNode
	doublyLinkedlist.tail = &newNode
	doublyLinkedlist.length++
	return
}

// Add a value to beginning of the list
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
	newNode.prev = doublyLinkedList.tail
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

func (doublyLinkedList *DoublyLinkedList) display() {
	if doublyLinkedList.length == 0 {
		fmt.Println("Linked list is empty")
		return
	}

	ptr := doublyLinkedList.head
	for i := 0; i < doublyLinkedList.length; i++ {
		fmt.Print(ptr.value, " ")
		ptr = ptr.next
	}
	fmt.Printf("\nLength: %d", doublyLinkedList.length)
	return
}

func main() {
	dll := DoublyLinkedList{}
	dll.append(10)
	dll.append(1)
	dll.append(0)
	dll.prepend(4)
	dll.prepend(3)
	dll.insert(3, 7)
	dll.display()
	fmt.Println(dll.head.next.next.value)
}
