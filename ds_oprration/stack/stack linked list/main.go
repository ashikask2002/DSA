package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Push(a int) {
	node := &Node{a, nil}

	if s.head == nil {
		s.head = node
		s.size++
	} else {
		node.next = s.head
		s.head = node
		s.size++
	}
}

func (s *Stack) Pop() {
	value := s.head.data
	s.head = s.head.next
	fmt.Println("popped value is", value)
}

func (s *Stack) Display() {
	temp := s.head

	for temp != nil {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
	}
}

func (s *Stack) Size() int {
	return s.size
}

func main() {
	s := Stack{}
	s.Push(3)
	s.Push(45)
	s.Push(5)
	s.Push(12)
	s.Push(34)
	s.Pop()
	s.Display()
	fmt.Println("size of stack is ", s.Size())

}
