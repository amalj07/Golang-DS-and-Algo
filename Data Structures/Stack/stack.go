package main

import "fmt"

var size = 10

type Stack struct {
	items []int
	top   int
}

func (stack *Stack) push(value int) {
	if stack.top == size {
		fmt.Println("Stack is full")
		return
	}

	stack.items = append(stack.items, value)
	stack.top++
	return
}

func (stack *Stack) pop() {
	if stack.top == 0 {
		fmt.Println("Stack is empty")
		return
	}

	stack.items = stack.items[:stack.top-1]
	stack.top--
}

func (stack *Stack) peek() {
	if stack.top == 0 {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println(stack.items[stack.top-1])
}

func (stack *Stack) print() {
	if stack.top == 0 {
		fmt.Println("Stack is empty")
		return
	}

	for i := 0; i < stack.top; i++ {
		fmt.Print(stack.items[i], " ")
	}
}

func main() {
	s := Stack{}
	s.push(4)
	s.push(2)
	s.push(9)
	s.push(3)
	s.peek()
	s.pop()
	s.print()
}
