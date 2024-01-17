package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
}

func (s *Stack) Pop() int {
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	return toremove
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func RemoveEvenNumbers(s *Stack) {
	if s.IsEmpty() {
		return
	}
	toremove := s.Pop()
	RemoveEvenNumbers(s)
	if toremove%2 != 0 {
		s.Push(toremove)
	}
}

func main() {
	l := &Stack{}

	l.Push(2)
	l.Push(5)
	l.Push(8)
	l.Push(1)
	l.Push(3)
	fmt.Println("array before ", l.data)
	RemoveEvenNumbers(l)
	fmt.Println("array after removal even numbers", l.data)
}
