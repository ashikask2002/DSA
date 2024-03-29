package main

import "fmt"

type Stack struct {
	data []int
	size int
}

func main() {
	l := &Stack{}
	l.Push(4)
	l.Push(1)
	l.Push(5)
	l.Push(2)
	l.Push(9)
	fmt.Println("array before", l.data)
	middleIndex := l.size / 2
	currentIndex := 0
	l.Removemiddle(middleIndex, currentIndex)
	fmt.Println("array after",l.data)
}

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
	s.size++
}

func (s *Stack) Pop() int {
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	return toremove
}

func (s *Stack) Removemiddle(middleIndex int, currentIndex int) int {
	if middleIndex == currentIndex {
		return s.Pop()
	}
	temp := s.Pop()
	mid := s.Removemiddle(middleIndex, currentIndex+1)
	s.Push(temp)

	return mid

}
