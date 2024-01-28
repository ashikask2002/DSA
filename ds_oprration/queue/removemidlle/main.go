package main

import "fmt"

type Queue struct {
	data []int
	size int
}

func main() {
	l := &Queue{}
	l.enqueue(4)
	l.enqueue(0)
	l.enqueue(1)
	l.enqueue(3)
	l.enqueue(5)
	fmt.Println("array before", l.data)
	middleIndex := l.size / 2
	currentIndex := 0
	l.Removemiddle(middleIndex, currentIndex)
	fmt.Println("array after", l.data)
}

func (q *Queue) enqueue(a int) {
	q.data = append(q.data, a)
	q.size++
}

func (q *Queue) dequeue() int {
	toremove := q.data[0]
	q.data = q.data[1:]
	return toremove
}

func (q *Queue) Removemiddle(middleIndex int, currentIndex int) int {
	if currentIndex == middleIndex {
		return q.dequeue()
	}
	temp := q.dequeue()
	mid := q.Removemiddle(middleIndex, currentIndex+1)

	q.enqueue(temp)
	return mid

}
