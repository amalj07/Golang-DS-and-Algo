package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type CircularDoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Add a value to end of the list
func (cicularDoublyLinkedList *CircularDoublyLinkedList) append(value int) {
	newNode := Node{}
	newNode.value = value
	if cicularDoublyLinkedList.head == nil {
		cicularDoublyLinkedList.head = &newNode
		cicularDoublyLinkedList.tail = &newNode
		cicularDoublyLinkedList.length++
		return
	}

	newNode.next = cicularDoublyLinkedList.head
	newNode.prev = cicularDoublyLinkedList.tail
	cicularDoublyLinkedList.head.prev = &newNode
	cicularDoublyLinkedList.tail.next = &newNode
	cicularDoublyLinkedList.tail = &newNode
	cicularDoublyLinkedList.length++
	return
}

// Add a value to beginning of the list
func (cicularDoublyLinkedList *CircularDoublyLinkedList) prepend(value int) {
	newNode := Node{}
	newNode.value = value
	if cicularDoublyLinkedList.head == nil {
		cicularDoublyLinkedList.head = &newNode
		cicularDoublyLinkedList.tail = &newNode
		cicularDoublyLinkedList.length++
		return
	}

	newNode.next = cicularDoublyLinkedList.head
	newNode.prev = cicularDoublyLinkedList.tail
	cicularDoublyLinkedList.head.prev = &newNode
	cicularDoublyLinkedList.tail.next = &newNode
	cicularDoublyLinkedList.head = &newNode
	cicularDoublyLinkedList.length++
	return
}

// Add a value to a specific position of the list
func (cicularDoublyLinkedList *CircularDoublyLinkedList) insert(position, value int) {
	if position < 0 {
		fmt.Println("Error: Invalid position")
		return
	}

	if position == 0 {
		cicularDoublyLinkedList.prepend(value)
		return
	}

	if position > cicularDoublyLinkedList.length {
		cicularDoublyLinkedList.append(value)
		return
	}

	ptr := cicularDoublyLinkedList.head
	for i := 0; i < position; i++ {
		if i == position-1 {
			newNode := Node{}
			newNode.value = value
			newNode.next = ptr
			newNode.prev = ptr.prev
			ptr.prev.next = &newNode
			ptr.prev = &newNode
			cicularDoublyLinkedList.length++
			return
		}
		ptr = ptr.next
	}
}

// Remove a value from beginning of the list
func (cicularDoublyLinkedList *CircularDoublyLinkedList) removeFromBeginning() {
	if cicularDoublyLinkedList.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	prevNode := cicularDoublyLinkedList.head.prev
	nextNode := cicularDoublyLinkedList.head.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
	cicularDoublyLinkedList.head = nextNode
	cicularDoublyLinkedList.length--
	return
}

// Remove a value from end of the list
func (cicularDoublyLinkedList *CircularDoublyLinkedList) removeFromEnd() {
	if cicularDoublyLinkedList.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	prevNode := cicularDoublyLinkedList.tail.prev
	nextNode := cicularDoublyLinkedList.tail.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
	cicularDoublyLinkedList.tail = prevNode
	cicularDoublyLinkedList.length--
	return
}

// Remove a value from a specific position
func (cicularDoublyLinkedList *CircularDoublyLinkedList) remove(position int) {
	if position < 0 || position > cicularDoublyLinkedList.length {
		fmt.Println("Error: Invalid position")
		return
	}

	ptr := cicularDoublyLinkedList.head

	for i := 0; i < position; i++ {
		if i == position-1 {
			prevNode := ptr.prev
			nextNode := ptr.next
			prevNode.next = nextNode
			nextNode.prev = prevNode
			cicularDoublyLinkedList.length--
			return
		}
		ptr = ptr.next
	}
}

func (cicularDoublyLinkedList *CircularDoublyLinkedList) display() {
	if cicularDoublyLinkedList.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	ptr := cicularDoublyLinkedList.head
	for ptr.next != cicularDoublyLinkedList.head {
		fmt.Print(ptr.value, " ")
		ptr = ptr.next
	}
	fmt.Print(ptr.value)
	fmt.Printf("\nLength: %d", cicularDoublyLinkedList.length)
	return
}

func main() {
	cdll := CircularDoublyLinkedList{}
	cdll.append(10)
	cdll.append(1)
	cdll.append(0)
	cdll.prepend(4)
	cdll.prepend(3)
	cdll.insert(10, 7)
	cdll.removeFromBeginning()
	cdll.removeFromEnd()
	cdll.remove(4)
	cdll.display()
}
