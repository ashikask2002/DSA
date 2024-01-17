package main

import "fmt"

type Queue struct {
	data []int
}

func main() {
l := Queue{}
l.enqueue(5)
l.enqueue(3)
l.enqueue(1)
l.enqueue(9)
l.enqueue(4)
l.dequeue()
l.display()
}

func (q *Queue) enqueue(a int) {
	q.data = append(q.data, a)
}

func (q *Queue) dequeue() {
	if len(q.data) == 0 {
		fmt.Println("queue is empty")
	}

	q.data = q.data[1:]
}

func (q *Queue) display() {
	if len(q.data) == 0 {
		fmt.Println("queue is empty")
	}

	fmt.Println(q.data)
}
