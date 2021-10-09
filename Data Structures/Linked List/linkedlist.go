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

// Add value to beginning of the list
func (linkedList *LinkedList) prepend(value int) {
	newNode := Node{}
	newNode.value = value

	if linkedList.head == nil {
		newNode.next = nil
	} else {
		newNode.next = linkedList.head
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

	if linkedList.head == nil {
		linkedList.head = &newNode
		linkedList.length++
		return
	}

	ptr := linkedList.head

	for ptr.next != nil {
		ptr = ptr.next
	}

	ptr.next = &newNode
	linkedList.length++
}

// Add value to specified index of Linked list
func (linkedList *LinkedList) insert(position, value int) {
	if position <= 0 {
		fmt.Println("Error: Invalid position")
		return
	}

	newNode := Node{}
	newNode.value = value

	if position >= linkedList.length {
		linkedList.append(value)
		return
	}

	if position == 1 {
		linkedList.prepend(value)
		return
	}

	ptr := linkedList.head
	for i := 0; i != position-1; {
		ptr = ptr.next
		i++
	}
	newNode.next = ptr.next
	ptr.next = &newNode
	linkedList.length++
}

// Remove a value from the list
func (linkedList *LinkedList) remove(position int) {
	if position > linkedList.length || position < 0 {
		fmt.Println("Error: Invalid position")
		return
	}

	if position == 0 {
		linkedList.removeFromBeginning()
		return
	}

	if position == linkedList.length {
		linkedList.removeFromEnd()
		return
	}

	ptr := linkedList.head
	for i := 0; i <= position-2; i++ {
		if i == position-2 {
			ptr.next = ptr.next.next
			linkedList.length--
			return
		}
		ptr = ptr.next
	}

}

// Remove value from beginning
func (linkedList *LinkedList) removeFromBeginning() {
	if linkedList.length == 0 {
		fmt.Println("Error: Linked List is empty")
		return
	}

	linkedList.head = linkedList.head.next
	linkedList.length--
	return
}

// Remove value from end
func (linkedList *LinkedList) removeFromEnd() {
	if linkedList.length == 0 {
		fmt.Println("Error: Linked List is empty")
		return
	}

	ptr := linkedList.head
	for ptr.next.next != nil {
		ptr = ptr.next
	}

	ptr.next = nil
	linkedList.length--
}

func (linkedList *LinkedList) display() {
	if linkedList.head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	ptr := linkedList.head
	for ptr.next != nil {
		fmt.Print(ptr.value, " ")
		ptr = ptr.next
	}
	fmt.Print(ptr.value)
	fmt.Printf("\nLength: %d\n", linkedList.length)
}

func main() {
	l := LinkedList{}
	l.append(10)
	l.prepend(4)
	l.insert(1, 12)
	l.insert(43, 31)
	l.append(21)
	l.remove(4)
	l.removeFromBeginning()
	l.removeFromEnd()
	l.display()
}
