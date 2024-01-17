package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	head *Node
	end  *Node
}

func main() {
l := Queue{}
l.enqueue(2)
l.enqueue(6)
l.enqueue(1)
l.enqueue(9)
l.enqueue(3)
l.dequeue()
l.display()
}

func (q *Queue) enqueue(a int) {
	node := &Node{a, nil}

	if q.head == nil {
		q.head = node
	} else {
		q.end.next = node
	}
	q.end = node

}

func (q *Queue) dequeue() {
	if q.head == nil {
		fmt.Println("queue is empty")
	}
	q.head = q.head.next

}
func (q *Queue) display() {
	if q.head == nil {
		fmt.Println("queue is empty")
	}
	for q.head != nil {
		fmt.Println(q.head.data)
		q.head = q.head.next
	}
}
