package main

import "fmt"

type queue struct {
	data []int
}

func main() {
	l := &stackwithqueues{}
	l.push(1)
	l.push(4)
	l.push(7)
	fmt.Println(l.pop())
	fmt.Println(l.pop())
	fmt.Println(l.pop())
}
func (q *queue) enqueue(a int) {
	q.data = append(q.data, a)
}
func (q *queue) dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	toremove := q.data[0]
	q.data = q.data[1:]
	return toremove
}

type stackwithqueues struct {
	queue1 queue
	queue2 queue
}

func (s *stackwithqueues) push(a int) {
	if len(s.queue2.data) == 0 {
		s.queue1.enqueue(a)
	} else {
		s.queue2.enqueue(a)
	}
}

func (s *stackwithqueues) pop() int {
	var noEmpty, empty *queue
	if len(s.queue1.data) == 0 {
		noEmpty, empty = &s.queue2, &s.queue1
	} else {
		noEmpty, empty = &s.queue1, &s.queue2
	}

	if len(noEmpty.data) == 0 {
		return -1
	}

	for len(noEmpty.data) > 1 {
		item := noEmpty.dequeue()
		empty.enqueue(item)
	}
	return noEmpty.dequeue()
}
