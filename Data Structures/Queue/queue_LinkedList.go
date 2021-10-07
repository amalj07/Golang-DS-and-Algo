package main

import "fmt"

type Node struct {
	value int
	next  *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (queue *Queue) enqueue(value int) {
	newNode := Node{}
	newNode.value = value

	if queue.front == nil {
		queue.front = &newNode
		queue.rear = &newNode
		return
	}

	queue.rear.next = &newNode
	queue.rear = &newNode
	return
}

func (queue *Queue) dequeue() {
	if queue.front == nil {
		fmt.Println("Queue is emtpy")
		return
	}

	if queue.front == queue.rear {
		queue.front = nil
		queue.rear = nil
		return
	}

	queue.front = queue.front.next
	return
}

func (queue *Queue) display() {
	if queue.front == nil {
		fmt.Println("Queue is empty")
		return
	}

	ptr := queue.front
	for ptr != queue.rear {
		fmt.Print(ptr.value, " ")
		ptr = ptr.next
	}

	fmt.Print(queue.rear.value)
	return
}

func main() {
	q := Queue{}
	q.enqueue(4)
	q.enqueue(2)
	q.enqueue(8)
	q.enqueue(3)
	q.enqueue(1)
	q.dequeue()
	q.dequeue()
	q.display()
}
