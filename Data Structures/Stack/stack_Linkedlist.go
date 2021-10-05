package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top    *Node
	length int
}

func (stack *Stack) push(value int) {
	newNode := Node{}
	newNode.value = value

	ptr := stack.top
	stack.top = &newNode
	stack.top.next = ptr
	stack.length++
}

func (stack *Stack) pop() {
	if stack.length == 0 {
		fmt.Println("Stack is emtpy")
		return
	}

	stack.top = stack.top.next
	stack.length--
	return
}

func (stack *Stack) peek() {
	if stack.length == 0 {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println(stack.top.value)
}

func (stack *Stack) print() {
	if stack.length == 0 {
		fmt.Println("Stack is empty")
		return
	}

	ptr := stack.top

	for i := 0; i < stack.length; i++ {
		fmt.Print(ptr.value, " ")
		ptr = ptr.next
	}
	fmt.Printf("\nSize: %d\n", stack.length)
}

func main() {
	s := Stack{}

	s.push(4)
	s.push(1)
	s.push(9)
	s.peek()
	s.print()
	s.pop()
	s.peek()
	s.print()
}
