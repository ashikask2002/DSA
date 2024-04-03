package main

import "fmt"

type stack struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func main() {
	l := &stack{}
	l.Push(5)
	l.Push(2)
	l.Push(1)
	l.Print()
	l.Reverse()
	l.Print()

}

func (s *stack) Push(a int) {
	node := &Node{a, nil}
	node.next = s.head
	s.head = node
}
func (s *stack) Pop() {
	s.head = s.head.next
}

func (s *stack) Reverse() {
	var prev, current, next *Node

	current = s.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	s.head = prev
}

func (s *stack) Print() {
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println()
}


