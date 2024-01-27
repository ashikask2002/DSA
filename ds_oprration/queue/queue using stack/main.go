package main

import "fmt"

type stack struct {
	data []int
}

func main() {
	queue := queuewithstack{}

	// Enqueue some elements
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	// Dequeue elements
	fmt.Println(queue.dequeue()) // Output: 1
	fmt.Println(queue.dequeue()) // Output: 2

	// Enqueue more elements
	queue.enqueue(4)
	queue.enqueue(5)

	// Dequeue the remaining elements
	fmt.Println(queue.dequeue()) // Output: 3
	fmt.Println(queue.dequeue()) // Output: 4
	fmt.Println(queue.dequeue()) // Output: 5
	fmt.Println(queue.dequeue()) // Output: -1 (indicating an empty queue)
}

func (s *stack) push(a int) {
	s.data = append(s.data, a)
}
func (s *stack) pop() int {
	if len(s.data) == 0 {
		return -1
	}
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	return toremove
}

type queuewithstack struct {
	enqueustack  stack
	dequeuestack stack
}

func (q *queuewithstack) enqueue(a int) {
	q.enqueustack.push(a)
}

func (q *queuewithstack) dequeue() int{
	if len(q.dequeuestack.data) == 0{
		for len(q.enqueustack.data) >0{
			q.dequeuestack.push(q.enqueustack.pop())
		}
	}
	if len(q.dequeuestack.data) == 0{
		return -1
	}
	return q.dequeuestack.pop()
}
